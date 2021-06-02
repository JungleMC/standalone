package protocol

type Protocol int

const (
	ProtocolHandshake Protocol = iota
	ProtocolStatus
	ProtocolLogin
	ProtocolPlay
)

func (p Protocol) String() string {
	switch p {
	case ProtocolHandshake:
		return "handshake"
	case ProtocolStatus:
		return "status"
	case ProtocolLogin:
		return "login"
	case ProtocolPlay:
		return "play"
	}
	return "unknown"
}
