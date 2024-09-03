package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/1layar/universe/internal/payment_service/app/appconfig"
	"github.com/1layar/universe/pkg/logger"
	"github.com/sony/gobreaker/v2"
	"go.uber.org/zap"
)

type IpaymuService struct {
	Cb     *gobreaker.CircuitBreaker[[]byte]
	config *appconfig.Config
	logger *zap.SugaredLogger
}

func NewPaymuService(
	cb *gobreaker.CircuitBreaker[[]byte],
	config *appconfig.Config,
) *IpaymuService {
	return &IpaymuService{
		Cb:     cb,
		config: config,
		logger: logger.GetLogger().With("service", "ipaymu"),
	}
}

func (s *IpaymuService) GenSignature(body []byte) (string, error) {
	//generate signature
	bodyHash := sha256.Sum256(body)
	bodyHashToString := hex.EncodeToString(bodyHash[:])
	stringToSign := "POST:" + s.config.IpaymuVa + ":" + strings.ToLower(string(bodyHashToString)) + ":" + s.config.IpaymuKey

	h := hmac.New(sha256.New, []byte(s.config.IpaymuKey))
	_, err := h.Write([]byte(stringToSign))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(h.Sum(nil))
	// end generate signatrure

	return signature, nil
}
