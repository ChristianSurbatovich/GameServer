package main

import (
	"math/rand"
)

type LootArea struct{
	id int16
	tier int16
	randomSeed float32
	looted bool
	lootList []int16
}


func (area *LootArea) generateLoot(itemList map[int16][]baseItem){
	area.looted = true
	area.lootList = make(map[int16]baseItem)
	listLength := len(itemList[area.tier])
	numItems := int16(rand.Intn(4) + 1)
	for i := int16(0); i < numItems; i++{
		area.lootList[i] = itemList[area.tier][rand.Intn(listLength)]
	}

}

func PopulateArea() LootArea{

}

func populateAreas()map[int16]LootArea {
	list := make(map[int16]LootArea)
	list[2000] = LootArea{2000,1,1,false,nil}
	list[2001] = LootArea{2001,1,1,false,nil}
	list[2002] = LootArea{2002,1,1,false,nil}
	list[2003] = LootArea{2003,2,1,false,nil}
	list[2004] = LootArea{2004,2,1,false,nil}
	list[2005] = LootArea{2005,2,1,false,nil}
	list[2006] = LootArea{2006,2,1,false,nil}
	list[2007] = LootArea{2007,3,1,false,nil}
	list[2008] = LootArea{2008,3,1,false,nil}
	list[2009] = LootArea{2009,3,1,false,nil}
	return list
}

func populateLoot() map[int16][]baseItem{
	list := make(map[int16][]baseItem)
	list[1] = make([]baseItem,3)
	list[1][0] = &genericStatItem{itemID:1100,statMods:[]statMod{{statID:CANNON_DAMAGE,value:0.5,percent:true,resource:false}}}
	list[1][1] = &genericStatItem{itemID:1101,statMods:[]statMod{{statID:MAX_HEALTH,value:50,percent:false,resource:true,resourceMode:PERCENT}}}
	list[1][2] = &genericStatItem{itemID:1102,statMods:[]statMod{{statID:SPEED,value:7,percent:false,resource:true,resourceMode:PERCENT}}}
	list[2] = make([]baseItem,3)
	list[2][0] = &genericStatItem{itemID:1100,statMods:[]statMod{{statID:CANNON_DAMAGE,value:0.75,percent:true,resource:false}}}
	list[2][1] = &genericStatItem{itemID:1101,statMods:[]statMod{{statID:MAX_HEALTH,value:100,percent:false,resource:true,resourceMode:PERCENT}}}
	list[2][2] = &genericStatItem{itemID:1102,statMods:[]statMod{{statID:SPEED,value:11,percent:false,resource:true,resourceMode:PERCENT}}}
	list[3] = make([]baseItem,3)
	list[3][0] = &genericStatItem{itemID:1100,statMods:[]statMod{{statID:CANNON_DAMAGE,value:1.0,percent:true,resource:false}}}
	list[3][1] = &genericStatItem{itemID:1101,statMods:[]statMod{{statID:MAX_HEALTH,value:150,percent:false,resource:true,resourceMode:PERCENT}}}
	list[3][2] = &genericStatItem{itemID:1102,statMods:[]statMod{{statID:SPEED,value:15,percent:false,resource:true,resourceMode:PERCENT}}}
	return list
}