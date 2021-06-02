package packet

type ClientboundStatusResponsePacket struct {
	Response string
}

type ClientboundStatusPongPacket struct {
	Time int64
}
