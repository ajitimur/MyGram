package domain

type Message struct {
	Message string `json:"message"`
}

type TokenMessage struct {
	Token string `json:"token"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
