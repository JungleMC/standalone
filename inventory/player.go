package inventory

import (
	"github.com/junglemc/JungleTree/item"
)

type Player struct {
	CraftingInput  [4]item.Slot
	CraftingOutput item.Slot
	Armor          [4]item.Slot
	Contents       [39]item.Slot
	OffHand        item.Slot
}
