package packets

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/junglemc/JungleTree/pkg/chat"
	"github.com/junglemc/JungleTree/pkg/codec"
)

type ClientboundLoginDisconnectPacket struct {
	Reason *chat.Message
}

func (p ClientboundLoginDisconnectPacket) MarshalMinecraft() ([]byte, error) {
	data, err := json.Marshal(&p.Reason)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Must be string
	return codec.WriteString(string(data)), nil
}

type ClientboundLoginEncryptionRequest struct {
	ServerId    string
	PublicKey   []byte ``
	VerifyToken []byte ``
}

type ClientboundLoginSuccess struct {
	Uuid     uuid.UUID `size:"infer"`
	Username string
}

type ClientboundLoginCompressionPacket struct {
	Threshold int32 `type:"varint"`
}

type ClientboundLoginPluginRequest struct {
	MessageId int32 `type:"varint"`
	Channel   string
	Data      []byte `size:"infer"`
}
