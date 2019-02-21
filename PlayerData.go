package main

import "sync"

type PlayerData struct{
	transform *playerTransform
	stats *playerStats
	items *ItemMap
	state *shipState
	lock *sync.RWMutex
}

func NewPlayerData() *PlayerData{
	var data PlayerData
	return &data
}
