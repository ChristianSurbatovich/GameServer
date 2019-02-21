package main

import (
	"sync"
	"encoding/binary"
)

type ItemMap struct{
	inventory map[int16]*inventorySlot
	lootAreas map[int16]lootArea
	bagSize int16
	filledSpaces int16
	bagStart int16
	bagEnd int16
	localID int16
	lock *sync.RWMutex
}

func NewItemMap() *ItemMap {
	var itemMap ItemMap
	itemMap.inventory = make(map[int16]*inventorySlot)
	itemMap.lootAreas = populateAreas()
	return &itemMap
}

func (itemMap *ItemMap) AddItemFromArea(lootArea int16, lootID int16, slotID int16){
	if !itemMap.inventory[slotID].open {

	}else{

	}
}

const(
	PICKUP_SUCCESS = iota
	PICKUP_FAIL_NO_ITEM = iota
	PICKUP_FAIL_NO_SPACE = iota
)


func (itemMap *ItemMap) PickupItem(areaID int16, itemID int16) (int16,int){
	item, exists := itemMap.lootAreas[areaID].lootList[itemID]
	if exists{
		for i := itemMap.bagStart ; i <= itemMap.bagEnd; i++ {
			if itemMap.inventory[i].open{
				itemMap.inventory[i].item = item
				delete(itemMap.lootAreas[areaID].lootList,itemID)
				return i,PICKUP_SUCCESS
			}
		}
		return  0,PICKUP_FAIL_NO_SPACE
	}
	return 0,PICKUP_FAIL_NO_ITEM
}
func (itemMap *ItemMap) SwapItems(source int16, destination int16){
	if itemMap.inventory[destination].open{
		// item can be moved
		itemMap.inventory[destination] = itemMap.inventory[source]
		itemMap.inventory[source].item = nil
	}else{
		// items have to be swapped
		temp := itemMap.inventory[destination]
		itemMap.inventory[destination] = itemMap.inventory[source]
		itemMap.inventory[source] = temp
	}
	itemMap.inventory[source].Refresh()
	itemMap.inventory[destination].Refresh()
}