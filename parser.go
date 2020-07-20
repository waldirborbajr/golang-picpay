package picpay

// StatusResultJSON returns status result
type StatusResultJSON struct {
	AuthorizationID string `json:"authorizationId,omitempty"`
	ReferenceID     string `json:"referenceId,omitempty"`
	Status          string `json:"status,omitempty"`
	Message         string `json:"message,omitempty"`
	Code            string `json:"code,omitempty"`
}

// PaymentResultJSON returns payment result
type PaymentResultJSON struct {
	ReferenceID string `json:"referenceId"`
	PaymentURL  string `json:"paymentUrl"`
	ExpiresAt   string `json:"expiresAt"`
	Qrcode      struct {
		Content string `json:"content"`
		Base64  string `json:"base64"`
	} `json:"qrcode"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Errors  []struct {
		Field   string `json:"field,omitempty"`
		Message string `json:"message,omitempty"`
	}
}
