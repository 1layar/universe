package service

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/1layar/universe/internal/payment_service/app/appconfig"
	"github.com/1layar/universe/pkg/logger"
	"github.com/imroc/req/v3"
	"github.com/sony/gobreaker/v2"
	"go.uber.org/zap"
)

type IakService struct {
	Cb        *gobreaker.CircuitBreaker[[]byte]
	config    *appconfig.Config
	logger    *zap.SugaredLogger
	operators []appconfig.IakOperatorPrefix
}

type SignedRequest struct {
	Username string `json:"username"`
	Sign     string `json:"sign"`
}

type CheckBalanceRequest struct {
	SignedRequest
}

type PriceListRequest struct {
	SignedRequest
	Status appconfig.IakProductStatus `json:"status"`
}

type DataResponse[T any] struct {
	Data T `json:"data"`
}

type CheckBalanceResponse struct {
	DataResponse[CheckBalanceData]
}

type Product struct {
	ProductCode        string                     `json:"product_code"`
	ProductDescription string                     `json:"product_description"`
	ProductNominal     string                     `json:"product_nominal"`
	ProductDetails     string                     `json:"product_details"`
	ProductPrice       float64                    `json:"product_price"`
	ProductType        string                     `json:"product_type"`
	ActivePeriod       string                     `json:"active_period"`
	Status             appconfig.IakProductStatus `json:"status"`
	IconURL            string                     `json:"icon_url"`
	ProductCategory    string                     `json:"product_category"`
}

type AccountPlan struct {
	Status       appconfig.IakPlnInqueryStatus `json:"status"`
	CustomerID   string                        `json:"customer_id"`
	MeterNo      string                        `json:"meter_no"`
	SubscriberID string                        `json:"subscriber_id"`
	Name         string                        `json:"name"`
	SegmentPower string                        `json:"segment_power"`
}

type PriceListData struct {
	PriceList []Product `json:"pricelist"`
	BaseData
}

type ProductListResponse struct {
	DataResponse[PriceListData]
}

type BaseData struct {
	Message string                  `json:"message"`
	Rc      appconfig.IakStatusCode `json:"rc"`
}

type CheckBalanceData struct {
	Balance int `json:"balance"`
	BaseData
}

type InqueryPlnData struct {
	AccountPlan
	BaseData
}

type InqueryPlnResponse struct {
	DataResponse[InqueryPlnData]
}

type InqueryPlnRequest struct {
	SignedRequest
	CustomerId string `json:"customer_id"`
}

type TopUpRequest struct {
	SignedRequest
	RefId       string `json:"ref_id"`
	CustomerId  string `json:"customer_id"`
	ProductCode string `json:"product_code"`
}

type TopupData struct {
	BaseData
	RefID       string `json:"ref_id"`
	Status      int64  `json:"status"`
	ProductCode string `json:"product_code"`
	CustomerID  string `json:"customer_id"`
	Price       int64  `json:"price"`
	Balance     int64  `json:"balance"`
	TrID        int64  `json:"tr_id"`
}

type TopUpResponse struct {
	DataResponse[TopupData]
}

type InqueryGameServerRequest struct {
	SignedRequest
	GameCode string `json:"game_code"`
}

type Server struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CheckStatusRequest struct {
	SignedRequest
	RefId string `json:"ref_id"`
}

type InqueryGameServerData struct {
	BaseData
	Status  appconfig.IakStatusCode `json:"status"`
	Servers []Server                `json:"servers"`
}

type InqueryGameServerResponse struct {
	DataResponse[InqueryGameServerData]
}

type CheckStatusData struct {
	BaseData
	RefID       string `json:"ref_id"`
	Status      int64  `json:"status"`
	ProductCode string `json:"product_code"`
	CustomerID  string `json:"customer_id"`
	Price       int64  `json:"price"`
	Sn          string `json:"sn"`
	Balance     int64  `json:"balance"`
	TrID        int64  `json:"tr_id"`
}

type CheckStatusResponse struct {
	DataResponse[CheckStatusData]
}

func NewIakService(
	cb *gobreaker.CircuitBreaker[[]byte],
	config *appconfig.Config,
) *IakService {
	return &IakService{
		Cb:     cb,
		config: config,
		logger: logger.GetLogger().With("service", "iak"),
		operators: []appconfig.IakOperatorPrefix{
			appconfig.AXIS,
			appconfig.ByU,
			appconfig.INDOSAT,
			appconfig.TELKOMSEL,
			appconfig.THREE,
			appconfig.XL,
			appconfig.ByU,
		},
	}
}

/**
 * Get signature by formula: md5({username}+{api_key}+{additional})
 * @param additional key for each api call
 */
func (s *IakService) GetSign(additional string) string {
	username := s.config.IakUsername
	apiKey := s.config.IakApiKey
	md5Sign := md5.Sum([]byte(username + apiKey + additional))
	return string(md5Sign[:])
}

// GetClient returns a req.Client with the "username" and "sign" headers set using the
// IAK API key and the given additional string.
func (s *IakService) GetClient(additional string) *req.Client {
	sign := s.GetSign(additional)

	return req.C().
		SetCommonHeader("username", s.config.IakUsername).
		SetCommonHeader("sign", string(sign[:]))
}

func (s *IakService) SendIak(path string, body any, additional string) ([]byte, error) {
	return s.Cb.Execute(func() ([]byte, error) {
		url := s.config.IakPrepaidUrl + path
		client := s.GetClient(additional)

		resp, err := client.R().SetBodyJsonMarshal(body).Post(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}

		if !resp.IsSuccessState() {
			// log the body
			s.logger.Error(resp.String())
			return nil, errors.New("failed to send iak request")
		}

		return body, nil
	})
}

/** API to get remaining balance in your IAK wallet. */
func (s *IakService) GetBalance() (*CheckBalanceResponse, error) {
	body, err := s.SendIak(
		"/api/check-balance",
		CheckBalanceRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign("bl"),
			},
		},
		"bl",
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data CheckBalanceResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

/** API to get pricelist of IAK prepaid products. */
func (s *IakService) GetPriceList(_type appconfig.IakProductType, operator appconfig.IakProductOperator, status appconfig.IakProductStatus) (*ProductListResponse, error) {
	body, err := s.SendIak(
		fmt.Sprintf("/api/pricelist/%s/%s", _type, operator),
		PriceListRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign("pl"),
			},
			Status: status,
		},
		"pl",
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data ProductListResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *IakService) CheckOperatorPrefix(phone string) (*appconfig.IakOperatorPrefix, error) {
	for _, op := range s.operators {
		for _, prefix := range op {
			if strings.HasPrefix(phone, prefix) {
				return &op, nil
			}
		}
	}

	return nil, errors.New("operator not found")
}

// API to check whether PLN Prepaid Subscriber is valid or invalid.
func (s *IakService) InquiryPLN(customerId string) (*InqueryPlnResponse, error) {
	body, err := s.SendIak(
		"/api/inquiry-pln",
		InqueryPlnRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign(customerId),
			},
			CustomerId: customerId,
		},
		customerId,
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data InqueryPlnResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// API to top up prepaid products.
func (s *IakService) TopUp(refId string, customerId string, productCode string) (*TopUpResponse, error) {
	body, err := s.SendIak(
		"/api/topup",
		TopUpRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign(refId),
			},
			RefId:       refId,
			CustomerId:  customerId,
			ProductCode: productCode,
		},
		refId,
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data TopUpResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// API to top up prepaid products.
// IAK will automatically detect the operator for each number in hp field that you send.
// Change the pulsa_code field to one of the code we listed below to activate this feature.
func (s *IakService) TopuPulsa(refId string, amount int64, hp string) (*TopUpResponse, error) {
	body, err := s.SendIak(
		"/api/topup",
		TopUpRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign(hp),
			},
			CustomerId:  hp,
			RefId:       refId,
			ProductCode: fmt.Sprintf("pulsa%d", amount),
		},
		hp,
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data TopUpResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// API to check game server list.
func (s *IakService) InquiryGameServer(gameCode appconfig.ServerListCode) (*InqueryGameServerResponse, error) {
	body, err := s.SendIak(
		"/api/inquiry-game-server",
		InqueryGameServerRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign(string(gameCode)),
			},
			GameCode: string(gameCode),
		},
		string(gameCode),
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data InqueryGameServerResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *IakService) GetGameFormat(gameCode appconfig.GameCode, info map[string]string) (string, error) {
	codeTemplate := appconfig.CODE_TEMPLATE[gameCode]

	for k, v := range info {
		codeTemplate = strings.ReplaceAll(codeTemplate, "{"+k+"}", v)
	}

	// when there { and } throw error
	if strings.Contains(codeTemplate, "{") || strings.Contains(codeTemplate, "}") {
		return "", errors.New("mising code parameter")
	}

	return codeTemplate, nil
}

// API to check status prepaid transaction.
func (s *IakService) CheckStatus(refId string) (*CheckStatusResponse, error) {
	body, err := s.SendIak(
		"/api/check-status",
		CheckStatusRequest{
			SignedRequest: SignedRequest{
				Username: s.config.IakUsername,
				Sign:     s.GetSign(refId),
			},
			RefId: refId,
		},
		refId,
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data CheckStatusResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
