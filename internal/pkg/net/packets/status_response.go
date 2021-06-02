package packets

import (
	"github.com/google/uuid"
	"github.com/junglemc/JungleTree/pkg/chat"
)

type ServerListResponse struct {
	Description *chat.Message     `json:"description,omitempty"`
	Players     ServerListPlayers `json:"players"`
	Version     GameVersion       `json:"version"`
	Favicon     string            `json:"favicon,omitempty"`
}

type GameVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type ServerListPlayers struct {
	Max    int                `json:"max"`
	Online int                `json:"online"`
	Sample []ServerListPlayer `json:"sample,omitempty"`
}

type ServerListPlayer struct {
	Name string    `json:"name"`
	Id   uuid.UUID `json:"id"`
}
