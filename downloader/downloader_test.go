package downloader_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/marcuxyz/golang-picpay/downloader"
	"github.com/marcuxyz/golang-picpay/response"

	picpay "github.com/marcuxyz/golang-picpay"
	"github.com/stretchr/testify/assert"
)

var status = response.StatusResponse{
	AuthorizationID: "123",
	ReferenceID:     "1203",
	Status:          "paid",
}

func TestMakeDownloaderStatusCode200(t *testing.T) {
	server := createServer()
	defer server.Close()

	URL, _ := url.Parse(server.URL)

	picpay := picpay.Picpay{
		Token: "",
		URL:   URL,
	}

	resp, err := downloader.MakeDownloader(http.MethodGet, server.URL, picpay.Token, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func createServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		result, _ := json.Marshal(&status)
		w.Write(result)
	})
	return httptest.NewServer(handler)
}
