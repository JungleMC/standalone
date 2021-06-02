package util

type Difficulty byte

const (
	Peaceful Difficulty = iota
	Easy
	Normal
	Hard
)

func (d Difficulty) String() string {
	switch d {
	case Peaceful:
		return "peaceful"
	case Easy:
		return "easy"
	case Normal:
		return "normal"
	case Hard:
		return "hard"
	}
	return ""
}

func DifficultyByName(name string) Difficulty {
	switch name {
	case "peaceful":
		return Peaceful
	case "easy":
		return Easy
	case "normal":
		return Normal
	case "hard":
		return Hard
	default:
		return Normal
	}
}
