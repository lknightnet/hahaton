package model

type Chat struct {
	Messages []Message `json:"messages"`
	Name     string    `json:"name"`
	SurName  string    `json:"surName"`
	MidName  string    `json:"midName"`
}

type Message struct {
	Text         string `json:"text"`
	FromOperator bool   `json:"fromOperator"`
}
