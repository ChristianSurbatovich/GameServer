package main

import (
	"bytes"
	"math/big"
	"encoding/binary"
	"go.opencensus.io/tag"
	"log"
)

type UpdateStage int8

const (
	UpdateFull UpdateStage = 0
	UpdateVelocity UpdateStage = 1
	UpdatePosition UpdateStage = 2
)

var UpdateOrder = []UpdateStage{UpdateFull,UpdateVelocity,UpdatePosition,UpdateVelocity}

type GameState struct{
	players map[int16]*PlayerData
	gameObjects map[int16]*GameObject
	messageQueue *Queue
	outgoingMessages [][]byte
	taggedPositions []TaggedMessage
	individualMessageQueue map[int16]*Queue
	slowUpdateRange float32
	noUpdateRange float32
	updateMessageStartSize int
}

func NewGameState() *GameState{
	return &GameState{players:make(map[int16]*PlayerData),gameObjects:make(map[int16]*GameObject),messageQueue:NewQueue(),outgoingMessageQueue:NewQueue()}
}

func (state *GameState) QueueMessage(message *RWMessage){
	state.messageQueue.Push(message)
}

func (state *GameState) QueueIndividualMessage(player int16, message []byte){
	state.individualMessageQueue[player].Push(message)
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
				playerID := message.ReadInt16()
				itemID := message.ReadInt16()
			case MOVE_ITEM:
				playerID := message.ReadInt16()
				itemSlot1 := message.ReadInt16()
				itemSlot2 := message.ReadInt16()
			case USE_ITEM:
				playerID := message.ReadInt16()
				itemSlotID := message.ReadInt16()
				}
			}
		}
	}
}

func (state *GameState) GenerateTaggedPositions(){
	state.taggedPositions = make([]TaggedMessage,len(state.players))
	for id,player := range state.players{
		message := RWMessage{new(bytes.Buffer)}
		message.Write(POSITION_FULL)
		message.Write(id)
		message.WriteVector(player.transform.position)
		message.WriteVector(player.transform.rotation)
		message.WriteVector(player.transform.velocity)
		message.WriteVector(player.transform.rotationVelocity)
		message.Write(player.transform.locationTime)
		state.taggedPositions = append(state.taggedPositions,TaggedMessage{player.transform.position,true,1000,message.Bytes()})
	}
}

func (state *GameState) GetUpdateMessage(playerID int16)[]byte{
	messageBuffer := new(bytes.Buffer)
	playerPosition := state.players[playerID].transform.position


	individualMessages := state.individualMessageQueue[playerID]
	for numMessages := individualMessages.Size(); numMessages > 0; numMessages--{
		tempMessage := individualMessages.Pop()
		if message, ok := tempMessage.([]byte); ok{
			messageBuffer.Write(message)
		}else{
			log.Println("Bad Message In Queue")
			log.Printf("PlayerID: %v\nMessage: %v\n",playerID,tempMessage)
		}
	}

	for _,message := range state.outgoingMessages{
		messageBuffer.Write(message)
	}
	for _,message := range state.taggedPositions{
		if message.CheckTag(playerPosition){
			messageBuffer.Write(message.bytes)
		}
	}

	return messageBuffer.Bytes()
}