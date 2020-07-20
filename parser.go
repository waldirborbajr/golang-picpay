package picpay

// StatusResultJSON returns status result
type StatusResultJSON struct {
	AuthorizationID string `json:"authorizationId"`
	ReferenceID     string `json:"referenceId"`
	Status          string `json:"status"`
	Message         string `json:"message"`
	Code            string `json:"code"`
}

type CancellationResultJSON struct {
	CancellationID string `json:"cancellationId"`
	ReferenceID    string `json:"referenceId"`
	Message        string `json:"message"`
	Code           string `json:"code"`
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
	Message string `json:"message"`
	Code    string `json:"code"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
}
