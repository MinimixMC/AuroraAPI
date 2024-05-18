package chat

type Message struct {
	Text *string `json:"text,omitempty"`

	Color string `json:"color,omitempty"`

	Bold          bool `json:"bold,omitempty,string"`
	Italic        bool `json:"italic,omitempty,string"`
	Underlined    bool `json:"underlined,omitempty,string"`
	Strikethrough bool `json:"strikethrough,omitempty,string"`
	Obfuscated    bool `json:"obfuscated,omitempty,string"`

	Extra []Message `json:"extra,omitempty"`

	Translate string    `json:"translate,omitempty"`
	With      []Message `json:"with,omitempty"`
}

func BuildMessage(t string) Message {
	message := Message{
		Text: &t,
	}
	return message
}
