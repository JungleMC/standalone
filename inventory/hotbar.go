package inventory

import (
	"github.com/junglemc/JungleTree/item"
)

type Hotbar struct {
	SelectedIndex int
	Slots         [9]*item.Slot
}

func NewHotbar() *Hotbar {
	result := &Hotbar{}
	for i := range result.Slots {
		result.Slots[i] = item.DefaultSlot()
	}
	return result
}
