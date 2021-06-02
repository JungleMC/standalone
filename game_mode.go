package JungleTree

type GameMode byte

const (
	Survival  GameMode = 0x00
	Creative           = 0x01
	Adventure          = 0x02
	Spectator          = 0x03
)

func (g GameMode) String() string {
	switch g {
	case Survival:
		return "survival"
	case Creative:
		return "creative"
	case Adventure:
		return "adventure"
	case Spectator:
		return "spectator"
	}
	return ""
}

func GameModeByName(name string) GameMode {
	switch name {
	case "survival":
		return Survival
	case "creative":
		return Creative
	case "adventure":
		return Adventure
	case "spectator":
		return Spectator
	default:
		return Survival
	}
}
