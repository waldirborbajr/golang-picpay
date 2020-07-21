package picpay

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/marcuxyz/golang-picpay/downloader"
	"github.com/marcuxyz/golang-picpay/response"
)

// Picpay base
type Picpay struct {
	Token string
	URL   *url.URL
}

// New Initializes an instance of the picpay object
func New(token string) *Picpay {
	URL, err := url.Parse("https://appws.picpay.com/ecommerce/public/payments/")
	if err != nil {
		log.Fatal(err)
	}

	return &Picpay{
		Token: token,
		URL:   URL,
	}
}

// GetOrderStatus get current status of order
func (p *Picpay) GetOrderStatus(referenceID string) (*response.StatusResponse, error) {
	p.URL.Path = path.Join(p.URL.Path, referenceID, "/status")

	resp, err := downloader.MakeDownloader(http.MethodGet, p.URL.String(), p.Token, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(response.StatusResponse)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// PayOrder create payment with Picpay
func (p *Picpay) PayOrder(buyer interface{}) (*response.PaymentResponse, error) {
	byte, err := json.Marshal(buyer)
	if err != nil {
		return nil, err
	}

	resp, err := downloader.MakeDownloader(http.MethodPost, p.URL.String(), p.Token, byte)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(response.PaymentResponse)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// CancelOrder cancel an order that has been paid!
func (p *Picpay) CancelOrder(authorizationID, referenceID string) (*response.CancellationResponse, error) {
	p.URL.Path = path.Join(p.URL.Path, referenceID, "/cancellations")

	body, err := json.Marshal(map[string]string{"authorizationId": authorizationID})
	if err != nil {
		return nil, err
	}

	resp, err := downloader.MakeDownloader(http.MethodPost, p.URL.String(), p.Token, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := new(response.CancellationResponse)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
