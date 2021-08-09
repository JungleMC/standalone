package chat

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/junglemc/JungleTree/pkg/codec"
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

func (c *Message) MarshalMinecraft() ([]byte, error) {
	data, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Must be string
	return codec.WriteString(string(data)), nil
}

func (c *Message) UnmarshalMinecraft(buf *bytes.Buffer) error {
	data := codec.ReadString(buf)
	err := json.Unmarshal([]byte(data), &buf)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
