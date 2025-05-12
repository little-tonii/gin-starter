package godoc

type MessageResponse struct {
	Message string `json:"message"`
}

type MessagesResponse struct {
	Messages []string `json:"message"`
}
