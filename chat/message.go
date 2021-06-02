package chat

import (
	"encoding/json"
	"log"
)

type Message struct {
	Bold          bool       `json:"bold,omitempty"`
	Italic        bool       `json:"italic,omitempty"`
	Underlined    bool       `json:"underlined,omitempty"`
	Strikethrough bool       `json:"strikethrough,omitempty"`
	Obfuscated    bool       `json:"obfuscated,omitempty"`
	Color         *ChatColor `json:"color,omitempty"`
	Insertion     string     `json:"insertion,omitempty"`
	ClickEvent    *Action    `json:"click_event,omitempty"`
	HoverEvent    *Action    `json:"hover_event,omitempty"`
	Text          string     `json:"text,omitempty"`
	Translate     string     `json:"translate,omitempty"`
	With          []*Message `json:"with,omitempty"`
}

func (c Message) String() string {
	data, err := json.Marshal(&c)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(data)
}
