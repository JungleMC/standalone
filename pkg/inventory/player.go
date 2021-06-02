package inventory

type Player struct {
	CraftingInput  [4]*Slot
	CraftingOutput *Slot
	Armor          [4]*Slot
	Contents       [39]*Slot
	OffHand        *Slot
}
