package ability

type PlayerAbilities byte

const (
	Invulnerable PlayerAbilities = 0x01 << iota
	Flying
	AllowFlying
	CreativeMode
)

func Set(abilities PlayerAbilities, ability ...PlayerAbilities) PlayerAbilities {
	for _, a := range ability {
		abilities = abilities | a
	}
	return abilities
}

func Unset(abilities PlayerAbilities, ability ...PlayerAbilities) PlayerAbilities {
	for _, a := range ability {
		abilities = abilities &^ a
	}
	return abilities
}

func (a PlayerAbilities) Has(ability PlayerAbilities) bool {
	return a&ability != 0
}
