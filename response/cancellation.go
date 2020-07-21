package response

// revive:disable

type CancellationResponse struct {
	CancellationID string `json:"cancellationId"`
	ReferenceID    string `json:"referenceId"`
	Message        string `json:"message"`
	Code           string `json:"code"`
}
