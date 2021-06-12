package packets

import (
	"github.com/google/uuid"
)

type ClientboundLoginDisconnectPacket struct {
	// Reason *chat.Message
	Reason string
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
