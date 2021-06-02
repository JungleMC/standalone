package protocol

type Protocol int

const (
	Handshake Protocol = iota
	Status
	Login
	Play
)

func (p Protocol) String() string {
	switch p {
	case Handshake:
		return "handshake"
	case Status:
		return "status"
	case Login:
		return "login"
	case Play:
		return "play"
	}
	return "unknown"
}
