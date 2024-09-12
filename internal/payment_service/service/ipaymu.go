package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/1layar/universe/internal/payment_service/app/appconfig"
	"github.com/1layar/universe/pkg/logger"
	"github.com/imroc/req/v3"
	"github.com/sony/gobreaker/v2"
	"go.uber.org/zap"
)

type DataResponse[T any] struct {
	Data T `json:"data"`
}

type IpaymuCheckBalanceResponse struct {
	IpaymuBaseResponse[IpaymuCheckBalance]
}

type IpaymuBaseResponse[T any] struct {
	DataResponse[T]
	Status  int64  `json:"Status"`
	Message string `json:"Message"`
}

type IpaymuCheckBalance struct {
	Va              string `json:"Va"`
	MerchantBalance int64  `json:"MerchantBalance"`
	MemberBalance   int64  `json:"MemberBalance"`
}

type IpaymuCheckTransactionResponse struct {
	Success bool `json:"Success"`
	IpaymuBaseResponse[CheckTransactionData]
}

type CheckTransactionData struct {
	TransactionID  int64     `json:"TransactionId"`
	SessionID      string    `json:"SessionId"`
	ReferenceID    any       `json:"ReferenceId"`
	RelatedID      int64     `json:"RelatedId"`
	Sender         string    `json:"Sender"`
	Receiver       string    `json:"Receiver"`
	Amount         int64     `json:"Amount"`
	Fee            int64     `json:"Fee"`
	Status         int64     `json:"Status"`
	StatusDesc     string    `json:"StatusDesc"`
	PaidStatus     string    `json:"PaidStatus"`
	IsLocked       bool      `json:"IsLocked"`
	Type           int64     `json:"Type"`
	TypeDesc       string    `json:"TypeDesc"`
	Notes          string    `json:"Notes"`
	CreatedDate    time.Time `json:"CreatedDate"`
	SuccessDate    time.Time `json:"SuccessDate"`
	ExpiredDate    time.Time `json:"ExpiredDate"`
	SettlementDate time.Time `json:"SettlementDate"`
	PaymentChannel string    `json:"PaymentChannel"`
	PaymentCode    string    `json:"PaymentCode"`
	BuyerName      string    `json:"BuyerName"`
	BuyerPhone     string    `json:"BuyerPhone"`
	BuyerEmail     string    `json:"BuyerEmail"`
}

type IpaymuHistoryPaymentResponse struct {
	Success bool `json:"Success"`
	IpaymuBaseResponse[IpaymuHistoryPaymentData]
}

type IpaymuHistoryPaymentData struct {
	Transaction []IpaymuTransaction `json:"Transaction"`
	Pagination  Pagination          `json:"Pagination"`
}

type Pagination struct {
	Total       int64 `json:"total"`
	Count       int64 `json:"count"`
	PerPage     int64 `json:"per_page"`
	CurrentPage int64 `json:"current_page"`
	TotalPages  int64 `json:"total_pages"`
}

type IpaymuTransaction struct {
	TransactionID  int64     `json:"TransactionId"`
	SessionID      *string   `json:"SessionId"`
	ReferenceID    *string   `json:"ReferenceId"`
	RelatedID      int64     `json:"RelatedId"`
	Sender         string    `json:"Sender"`
	Receiver       string    `json:"Receiver"`
	Amount         int64     `json:"Amount"`
	Fee            int64     `json:"Fee"`
	Status         int64     `json:"Status"`
	StatusDesc     string    `json:"StatusDesc"`
	PaidStatus     string    `json:"PaidStatus"`
	Type           int64     `json:"Type"`
	TypeDesc       string    `json:"TypeDesc"`
	Notes          *string   `json:"Notes"`
	IsEscrow       bool      `json:"IsEscrow"`
	CreatedDate    time.Time `json:"CreatedDate"`
	ExpiredDate    time.Time `json:"ExpiredDate"`
	SuccessDate    any       `json:"SuccessDate"`
	SettlementDate any       `json:"SettlementDate"`
	PaymentChannel string    `json:"PaymentChannel"`
	PaymentCode    string    `json:"PaymentCode"`
	BuyerName      string    `json:"BuyerName"`
	BuyerPhone     string    `json:"BuyerPhone"`
	BuyerEmail     string    `json:"BuyerEmail"`
}

type IpaymuPaymentMethodResponse struct {
	Status  int64   `json:"Status"`
	Success bool    `json:"Success"`
	Message string  `json:"Message"`
	Data    []Datum `json:"Data"`
}

type Datum struct {
	Code        string    `json:"Code"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Channels    []Channel `json:"Channels,omitempty"`
}

type Channel struct {
	Code                   string         `json:"Code"`
	Name                   string         `json:"Name"`
	Description            string         `json:"Description"`
	Logo                   string         `json:"Logo"`
	PaymentInstructionsDoc string         `json:"PaymentInstructionsDoc"`
	FeatureStatus          string         `json:"FeatureStatus"`
	HealthStatus           string         `json:"HealthStatus"`
	TransactionFee         TransactionFee `json:"TransactionFee"`
}

type TransactionFee struct {
	ActualFee     float64 `json:"ActualFee"`
	ActualFeeType string  `json:"ActualFeeType"`
	AdditionalFee int64   `json:"AdditionalFee"`
}

type IpaymuDirectPaymentRequest struct {
	Name            string     `json:"name"`
	Phone           string     `json:"phone"`
	Email           string     `json:"email"`
	Amount          int64      `json:"amount"`
	NotifyUrl       string     `json:"notifyUrl"`
	Expired         int64      `json:"expired"`
	Comments        *string    `json:"comments,omitempty"`
	ReferenceID     *string    `json:"referenceId,omitempty"`
	PaymentMethod   string     `json:"paymentMethod"`
	PaymentChannel  string     `json:"paymentChannel"`
	Product         *[]string  `json:"product,omitempty"`
	Qty             *[]int64   `json:"qty,omitempty"`
	Price           *[]int64   `json:"price,omitempty"`
	Weight          *[]float64 `json:"weight,omitempty"`
	Width           *[]float64 `json:"width,omitempty"`
	Height          *[]float64 `json:"height,omitempty"`
	Length          *[]float64 `json:"length,omitempty"`
	FeeDirection    string     `json:"feeDirection"`
	Escrow          *int8      `json:"escrow"`
	Account         *string    `json:"account"`
	DeliveryArea    *string    `json:"deliveryArea,omitempty"`
	DeliveryAddress *string    `json:"deliveryAddress,omitempty"`
	Shipping        *string    `json:"shipping,omitempty"`
	ShippingService *string    `json:"shippingService,omitempty"`
	PickupArea      *string    `json:"pickupArea,omitempty"`
}

type IpaymuDirectPaymentResponse struct {
	IpaymuBaseResponse[IpaymuDirectPaymentData]
}

type IpaymuDirectPaymentData struct {
	SessionID     string    `json:"SessionId"`
	TransactionID int64     `json:"TransactionId"`
	ReferenceID   string    `json:"ReferenceId"`
	Via           string    `json:"Via"`
	Channel       string    `json:"Channel"`
	PaymentNo     string    `json:"PaymentNo"`
	PaymentName   string    `json:"PaymentName"`
	Total         int64     `json:"Total"`
	Fee           int64     `json:"Fee"`
	Expired       time.Time `json:"Expired"`
}

type IpaymuService struct {
	Cb     *gobreaker.CircuitBreaker[[]byte]
	config *appconfig.Config
	logger *zap.SugaredLogger
}

func NewIPaymuService(
	cb *gobreaker.CircuitBreaker[[]byte],
	config *appconfig.Config,
) *IpaymuService {
	return &IpaymuService{
		Cb:     cb,
		config: config,
		logger: logger.GetLogger().With("service", "ipaymu"),
	}
}

func (s *IpaymuService) GenSignature(method string, body []byte) (string, error) {
	//generate signature
	bodyHash := sha256.Sum256(body)
	bodyHashToString := hex.EncodeToString(bodyHash[:])
	stringToSign := strings.ToUpper(method) + ":" + s.config.IpaymuVa + ":" + strings.ToLower(string(bodyHashToString)) + ":" + s.config.IpaymuKey

	h := hmac.New(sha256.New, []byte(s.config.IpaymuKey))
	_, err := h.Write([]byte(stringToSign))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(h.Sum(nil))
	// end generate signatrure

	return signature, nil
}

func (s *IpaymuService) GetClient(signature string) *req.Client {
	timetsamp := fmt.Sprintf("%d", time.Now().Unix())

	return req.C().
		SetCommonHeader("Content-Type", "application/json").
		SetCommonHeader("signature", signature).
		SetCommonHeader("va", s.config.IpaymuVa).
		SetCommonHeader("timestamp", timetsamp)
}

func (s *IpaymuService) SendIpaymuPost(path string, body any) ([]byte, error) {
	return s.Cb.Execute(func() ([]byte, error) {
		url := s.config.IpaymuBaseUrl + path
		bodySign, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}
		signature, err := s.GenSignature("POST", bodySign)

		if err != nil {
			return nil, err
		}

		client := s.GetClient(string(signature))

		resp, err := client.R().SetBody(bodySign).Post(url)
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
			return nil, errors.New("failed to send ipaymu request")
		}

		return body, nil
	})
}

func (s *IpaymuService) SendIpaymuGet(path string, body any) ([]byte, error) {
	return s.Cb.Execute(func() ([]byte, error) {
		url := s.config.IpaymuBaseUrl + path
		bodySign, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}
		signature, err := s.GenSignature("GET", []byte(""))

		if err != nil {
			return nil, err
		}

		client := s.GetClient(string(signature))

		resp, err := client.R().SetBody(bodySign).Get(url)
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
			return nil, errors.New("failed to send ipaymu get request")
		}

		return body, nil
	})
}

func (s *IpaymuService) CheckBalance() (*IpaymuCheckBalanceResponse, error) {
	body, err := s.SendIpaymuPost(
		"/api/v2/balance",
		map[string]string{
			"account": s.config.IpaymuVa,
		},
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data IpaymuCheckBalanceResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *IpaymuService) CheckTransaction(transactionId string) (*IpaymuCheckTransactionResponse, error) {
	body, err := s.SendIpaymuPost(
		"/api/v2/transaction/"+transactionId,
		map[string]string{
			"account": s.config.IpaymuVa,
		},
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data IpaymuCheckTransactionResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *IpaymuService) HistoryTransaction() (*IpaymuHistoryPaymentResponse, error) {
	body, err := s.SendIpaymuPost(
		"/api/v2/history",
		map[string]string{
			"orderBy": "id",
			"order":   "DESC",
			"limit":   "10",
		},
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data IpaymuHistoryPaymentResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *IpaymuService) IpaymuPaymentMethod() (*IpaymuPaymentMethodResponse, error) {
	body, err := s.SendIpaymuGet(
		"/api/v2/payment-channels",
		map[string]string{
			"account": s.config.IpaymuVa,
		},
	)

	if err != nil {
		return nil, err
	}

	// unmarshal response
	var data IpaymuPaymentMethodResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
