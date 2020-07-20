package picpay

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var status = StatusResultJSON{
	AuthorizationID: "123",
	ReferenceID:     "1203",
	Status:          "paid",
}

func TestMakeDownloaderStatusCode200(t *testing.T) {
	server := createServer()
	defer server.Close()

	picpay := Picpay{
		Token:   "",
		BaseURL: server.URL,
	}

	resp, err := MakeDownloader(http.MethodGet, server.URL, picpay.Token, nil)
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
