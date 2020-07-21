package response

// StatusResponse deliver the result of the request made by GetOrderStatus
type StatusResponse struct {
	AuthorizationID string `json:"authorizationId"`
	ReferenceID     string `json:"referenceId"`
	Status          string `json:"status"`
	Message         string `json:"message"`
	Code            string `json:"code"`
}
