package item

type Slot struct {
	Present bool
	ID      int32       `type:"varint" optional:"Present"`
	Count   byte        `optional:"Present"`
	Data    interface{} `type:"nbt" optional:"Present"`
}

func DefaultSlot() *Slot {
	return &Slot{
	}
}
