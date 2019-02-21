package main

import "sync"

type GameState struct{
	Players map[int16]*PlayerData
	MessageQueue *Queue
}

func NewGameState() *GameState{
	return &GameState{make(map[int16]*PlayerData),NewQueue()}
}

func (state *GameState) QueueMessage(message Message){
	state.MessageQueue.Push(message)
}
