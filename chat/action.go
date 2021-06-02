package chat

type ActionType string

type Action struct {
	Action *ActionType `json:"action,omitempty"`
	Value  string      `json:"value,omitempty"`
}
