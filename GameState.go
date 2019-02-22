package main


type GameState struct{
	players map[int16]*PlayerData
	staticObjects map[int16]
	messageQueue *Queue
	outgoingMessageQueue *Queue
}

func NewGameState() *GameState{
	return &GameState{make(map[int16]*PlayerData),NewQueue(),NewQueue()}
}

func (state *GameState) QueueMessage(message Message){
	state.messageQueue.Push(message)
}

func (state *GameState) ProcessMessages(){

}