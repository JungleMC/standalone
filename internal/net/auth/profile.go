package auth

import "github.com/google/uuid"

type Profile struct {
	ID         uuid.UUID         `json:"id"`
	Name       string            `json:"name"`
	Properties []ProfileProperty `json:"properties"`
}

type ProfileProperty struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signed    bool   `json:"-"`
	Signature string `json:"signature"`
}
