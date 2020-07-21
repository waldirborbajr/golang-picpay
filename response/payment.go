package response

// revive:disable

type PaymentResponse struct {
	ReferenceID string `json:"referenceId"`
	PaymentURL  string `json:"paymentUrl"`
	ExpiresAt   string `json:"expiresAt"`
	Qrcode      struct {
		Content string `json:"content"`
		Base64  string `json:"base64"`
	} `json:"qrcode"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
}
