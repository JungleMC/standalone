package inventory

type Hotbar struct {
	SelectedIndex int
	Slots         [9]*Slot
}

func NewHotbar() *Hotbar {
	result := &Hotbar{}
	for i := range result.Slots {
		result.Slots[i] = DefaultSlot()
	}
	return result
}
