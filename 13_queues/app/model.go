package app

type PushRequest struct {
	Message string `json:"message"`
}

type PopResponse struct {
	Message string `json:"message"`
}