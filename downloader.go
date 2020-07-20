package picpay

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MakeDownloader make request for consult picpay api and ensure value
func MakeDownloader(method string, URL string, token string, body []byte) ([]byte, error) {
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

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
