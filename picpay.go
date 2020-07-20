package picpay

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

// Picpay is a struct
// This object must not be instantiated manually, for this use the New() function
type Picpay struct {
	Token   string
	BaseURL string
}

// New Initializes an instance of the picpay object
func New(token string) *Picpay {
	return &Picpay{
		Token:   token,
		BaseURL: "https://appws.picpay.com/ecommerce/public/payments/",
	}
}

// GetOrderStatus get current status of order
func (p *Picpay) GetOrderStatus(referenceID string) (*StatusResultJSON, error) {
	URL, err := url.Parse(p.BaseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = path.Join(URL.Path, referenceID, "/status")

	resp, err := MakeDownloader(http.MethodGet, URL.String(), p.Token, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(StatusResultJSON)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// PayOrder create payment with Picpay
func (p *Picpay) PayOrder(buyer interface{}) (*PaymentResultJSON, error) {
	URL, err := url.Parse(p.BaseURL)
	if err != nil {
		return nil, err
	}

	byte, err := json.Marshal(buyer)
	if err != nil {
		return nil, err
	}

	resp, err := MakeDownloader(http.MethodPost, URL.String(), p.Token, byte)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(PaymentResultJSON)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *Picpay) CancelOrder(authorizationID, referenceID string) (*CancellationResultJSON, error) {
	URL, err := url.Parse(p.BaseURL)
	if err != nil {
		return nil, err
	}

	URL.Path = path.Join(URL.Path, referenceID, "/cancellations")

	body, err := json.Marshal(map[string]string{"authorizationId": authorizationID})
	if err != nil {
		return nil, err
	}

	resp, err := MakeDownloader(http.MethodPost, URL.String(), p.Token, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(CancellationResultJSON)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
