package downloader

import (
	"bytes"
	"net/http"
)

// MakeDownloader make request for consult picpay api and ensure value
func MakeDownloader(method string, URL string, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-picpay-token", token)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
