package model

type ChatHistory struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type ChatData struct {
	ChatHistory []ChatHistory `json:"chat_history"`
	Message     string        `json:"message"`
	Stream      bool          `json:"stream"`
	Model       string        `json:"model"`
	Connectors  []Connector   `json:"connectors,omitempty"`
}

type Connector struct {
	ID string `json:"id"`
}
