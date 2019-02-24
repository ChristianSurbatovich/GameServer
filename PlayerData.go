package main

import "sync"

type PlayerData struct{
	transform *Transform
	stats *PlayerStats
	items *ItemMap
	state *ShipData
	lock *sync.RWMutex
}

func NewPlayerData() *PlayerData{
	var data PlayerData
	return &data
}
