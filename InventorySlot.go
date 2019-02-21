package main

const(
	INVENTORY = iota
	HIDDEN = iota
	EQUIPPED = iota
	PASSIVE = iota
)

type inventorySlot struct{
	slotID int16
	item baseItem
	equipped bool
	slotType int
	open bool

}

func (slot *inventorySlot) Refresh(){

}
