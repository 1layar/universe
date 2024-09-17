package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/1layar/universe/internal/auth_service/internal/app/appconfig"
	"github.com/1layar/universe/internal/auth_service/model"
	"github.com/1layar/universe/internal/auth_service/model/dto"
	"github.com/1layar/universe/pkg/shared/command"
	"github.com/1layar/universe/pkg/shared/constant"
	"github.com/1layar/universe/pkg/shared/transport"
	"github.com/1layar/universe/pkg/shared/utils"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticService struct {
	CommandBus *cqrs.CommandBus
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
	Conf       *appconfig.Config
	Session    *SessionService
}

func NewAuthService(
	CommandBus *cqrs.CommandBus,
	Publisher *amqp.Publisher,
	Subscriber *amqp.Subscriber,
	Conf *appconfig.Config,
	session *SessionService,
) *AuthenticService {
	return &AuthenticService{
		CommandBus: CommandBus,
		Publisher:  Publisher,
		Subscriber: Subscriber,
		Conf:       Conf,
		Session:    session,
	}
}

func (s *AuthenticService) Login(ctx context.Context, cmdArgs command.LoginCommand) (*command.LoginResult, error) {
	replyCh, cancel, err := transport.GetReqRep[command.GetUserResult](
		constant.GET_USER_CMD,
		s.Publisher,
		s.Subscriber,
		ctx,
		s.CommandBus,
		&command.GetUserCommand{
			Email: cmdArgs.Email,
		},
	)

	if err != nil {
		return nil, err
	}

	defer cancel()

	select {
	case reply := <-replyCh:
		if err := reply.Error; err != nil {
			return nil, err
		}
		data := reply.HandlerResult
		valid := utils.CompPassword(data.Password, cmdArgs.Password)

		if !valid {
			return nil, errors.New("invalid_credential")
		}
		tokenPair, err := s.getPairToken(data)
		if err != nil {
			return nil, err
		}

		// save session
		_, err = s.Session.CreateSession(ctx, dto.CreateSessionDto{
			UserID:    data.Id,
			IP:        cmdArgs.Ip,
			UserAgent: cmdArgs.Ua,
			Retry:     0,
			Kind:      model.LoginKind,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to create session: %w", err)
		}

		return tokenPair, nil
	case <-time.After(time.Second * time.Duration(5)):
		return nil, errors.New("gateway_timeout")
	}
}

func (s *AuthenticService) getPairToken(data command.GetUserResult) (*command.LoginResult, error) {

	expDuration, err := time.ParseDuration(s.Conf.JwtExpTime)

	if err != nil {
		return nil, fmt.Errorf("get_exp_failed: %w", err)
	}

	accessExpTime := time.Now().Add(expDuration)
	refershExpTime := time.Now().Add(24 * time.Hour)

	// generate accessToken here
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(data.Id),
		ExpiresAt: jwt.NewNumericDate(accessExpTime),
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.Conf.JwtSecret))

	if err != nil {
		return nil, fmt.Errorf("sign_failed: %w", err)
	}

	// generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(data.Id),
		ExpiresAt: jwt.NewNumericDate(refershExpTime),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.Conf.JwtSecret))

	if err != nil {
		return nil, fmt.Errorf("sign_failed: %w", err)
	}

	return &command.LoginResult{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiredAt:    accessExpTime.Format(time.RFC822),
		Permission:   []string{},
	}, nil
}

func (s *AuthenticService) GetTokenPairByRefresh(refreshToken string) (*command.LoginResult, error) {
	panic("not implements")
}

func (s *AuthenticService) Register(ctx context.Context, cmd command.RegisterCommand) (*command.RegisterResult, error) {
	replyCh, cancel, err := transport.GetReqRep[command.AddUserResult](
		constant.ADD_USER_CMD,
		s.Publisher,
		s.Subscriber,
		ctx,
		s.CommandBus,
		&command.AddUserCommand{
			Email:    cmd.Email,
			Password: cmd.Password,
			Username: cmd.Username,
			Role:     1,
		},
	)

	if err != nil {
		return nil, err
	}

	defer cancel()
	select {
	case reply := <-replyCh:
		if err := reply.Error; err != nil {
			return nil, err
		}
		data := reply.HandlerResult
		tokenPair, err := s.getPairToken(command.GetUserResult{
			Id:    data.Id,
			Name:  data.Name,
			Email: data.Email,
		})

		if err != nil {
			return nil, err
		}

		// save session
		_, err = s.Session.CreateSession(ctx, dto.CreateSessionDto{
			UserID:    data.Id,
			IP:        cmd.Ip,
			UserAgent: cmd.Ua,
			Retry:     0,
			Kind:      model.LoginKind,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to create session: %w", err)
		}

		return &command.RegisterResult{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
			ExpiredAt:    tokenPair.ExpiredAt,
			Permission:   tokenPair.Permission,
		}, nil

	case <-time.After(time.Second * time.Duration(5)):
		return nil, errors.New("gateway_timeout")
	}
}

func (s *AuthenticService) ForgotPassword() {

}

func (s *AuthenticService) JwtToUser(ctx context.Context, cmd command.JwtToUserCommand) (*command.JwtToUserResponse, error) {
	token, err := jwt.Parse(cmd.Jwt, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(s.Conf.JwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		issuer, err := claims.GetSubject()

		if err != nil {
			return nil, fmt.Errorf("failed to get issuer: %w", err)
		}

		id, err := strconv.ParseInt(issuer, 10, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to parse jwt: %w", err)
		}

		replyCh, cancel, err := transport.GetReqRep[command.GetUserResult](
			constant.GET_USER_CMD,
			s.Publisher,
			s.Subscriber,
			ctx,
			s.CommandBus,
			&command.GetUserCommand{
				ID: int(id),
			},
		)
		if err != nil {
			return nil, err
		}

		defer cancel()

		select {
		case reply := <-replyCh:
			if err := reply.Error; err != nil {
				return nil, err
			}
			data := reply.HandlerResult
			return &command.JwtToUserResponse{
				User: data,
			}, nil

		case <-time.After(time.Second * time.Duration(15)):
			return nil, errors.New("gateway_timeout")
		}
	}

	return nil, errors.New("invalid claim")
}

func (s *AuthenticService) UserToJwt() {
	panic("not implements")
}
