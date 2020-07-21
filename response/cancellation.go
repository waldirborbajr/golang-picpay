package response

// CancellationResponse deliver the result of the request made by CancelOrder
type CancellationResponse struct {
	CancellationID string `json:"cancellationId"`
	ReferenceID    string `json:"referenceId"`
	Message        string `json:"message"`
	Code           string `json:"code"`
}
