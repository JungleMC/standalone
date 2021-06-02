package packets

type ServerboundLoginStartPacket struct {
	Username string
}

type ServerboundLoginEncryptionResponsePacket struct {
	SharedSecret []byte ``
	VerifyToken  []byte ``
}

type ServerboundLoginPluginResponsePacket struct {
	MessageID  int32 `type:"varint"`
	Successful bool
	Data       []byte `size:"infer" optional:"Successful"`
}
