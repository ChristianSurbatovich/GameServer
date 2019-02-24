package main


type GameState struct{
	players map[int16]*PlayerData
	gameObjects map[int16]*GameObject
	messageQueue *Queue
	outgoingMessageQueue *Queue
	individualMessageQueue map[int16]*Queue
}

func NewGameState() *GameState{
	return &GameState{players:make(map[int16]*PlayerData),gameObjects:make(map[int16]*GameObject),messageQueue:NewQueue(),outgoingMessageQueue:NewQueue()}
}

func (state *GameState) QueueMessage(message *RWMessage){
	state.messageQueue.Push(message)
}

func (state *GameState) ProcessMessages(){

	for i := state.messageQueue.Size(); i > 0; i--{
		message := state.messageQueue.Pop().(RWMessage)
		for message.UnreadData() {
			switch message.ReadByte() {
			case POSITION:
				playerTransform := state.players[message.ReadInt16()].transform
				playerTransform.position = message.ReadVector()
				playerTransform.rotation = message.ReadVector()
				playerTransform.locationTime = message.ReadFloat32()
			case VELOCITY:
				playerTransform := state.players[message.ReadInt16()].transform
				playerTransform.velocity = message.ReadVector()
				playerTransform.rotationVelocity = message.ReadVector()
			case POSITION_FULL:
				playerTransform := state.players[message.ReadInt16()].transform
				playerTransform.position = message.ReadVector()
				playerTransform.rotation = message.ReadVector()
				playerTransform.velocity = message.ReadVector()
				playerTransform.rotationVelocity = message.ReadVector()
				playerTransform.locationTime = message.ReadFloat32()
			case LOOT_ITEM:
				playerID := message.ReadInt16()
				state.individualMessageQueue[playerID].Push(state.players[playerID].items.AddItemFromArea(message.ReadInt16(),message.ReadInt16(),message.ReadInt16()))
			case LOOT_AREA:
				playerID := message.ReadInt16()
				state.individualMessageQueue[playerID].Push(state.players[playerID].items.LootArea(message.ReadInt16()))
			case PICKUP_ITEM:
			case MOVE_ITEM:
			case USE_ITEM:
			}
		}
	}
}