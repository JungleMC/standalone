package blocks

type BlockState struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	NumValues uint     `json:"num_values"`
	Values    []string `json:"values"`
}
