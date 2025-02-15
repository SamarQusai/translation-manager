package model

type DialogueRequest struct {
	Dialogue []Dialogue `json:"dialogue" validate:"required,validateConversations"`
}

type Dialogue struct {
	Speaker  string `json:"speaker" validate:"required"`
	Time     string `json:"time" validate:"required"`
	Sentence string `json:"sentence" validate:"required"`
}
