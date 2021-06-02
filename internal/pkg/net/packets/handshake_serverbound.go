package packets

type ServerboundHandshakePacket struct {
	ProtocolVersion int32 `type:"varint"`
	ServerHost      string
	ServerPort      uint16
	NextState       int32 `type:"varint"`
}

type ServerboundHandshakeLegacyPingPacket struct {
	Payload         byte
	ProtocolVersion byte
	Hostname        string
	Port            int16
}
