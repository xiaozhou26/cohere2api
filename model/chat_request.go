package model

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature"`
	PresencePenalty  float64   `json:"presence_penalty"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	TopP             float64   `json:"top_p"`
	Stream           bool      `json:"stream"`
	Model            string    `json:"model"`
}
