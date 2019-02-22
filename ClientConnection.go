package main

import (
	"net"
	"sync"
	"io"
	"encoding/binary"
	"bytes"
	"bufio"
	"log"
	"time"
)

type ClientConnection struct{
	ID int16
	conn net.Conn
	lock *sync.RWMutex
	reader *bufio.Reader
	gameState *GameState
	lengthBuffer []byte
	messageLength int
}

func NewClientConnection(id int16, newConnection net.Conn, state *GameState) *ClientConnection{
	return &ClientConnection{
		id,
		newConnection,
		&sync.RWMutex{},
		bufio.NewReader(newConnection),
		state,
		make([]byte,PREFIX_SIZE),
		0,
	}
}

func (client *ClientConnection) HandleClient(){
	for{
		n, err := io.ReadFull(client.reader,client.lengthBuffer)

		if err != nil {
			log.Printf("Bytes read: %d\n", n)
			log.Println(err)
			break
		}
		binary.Read(bytes.NewReader(client.lengthBuffer),binary.LittleEndian,&client.messageLength)
		messageBytes := make([]byte,client.messageLength)
		n, err = client.reader.Read(messageBytes)
		if err != nil{
			log.Println(err)
			if err == io.EOF{
				continue
			}else{
				break
			}
		}
		message := NewMessageReader(messageBytes)
		for message.UnreadData() == true {
			messageValue := message.ReadByte()
			switch messageValue{
			case SYNC:
				syncMessage := NewMessageWriter()
				syncMessage.WriteByte(SYNC)
				syncMessage.Write(float32(time.Now().Sub(serverStartTime).Seconds()))
				client.Send(syncMessage.Bytes())
			case POSITION:
				client.gameState.QueueMessage(&PositionMessage{client.ID,message.ReadVector(),message.ReadVector()})
			case VELOCITY:
				client.gameState.QueueMessage(&VelocityMessage{client.ID,message.ReadVector(),message.ReadVector()})
			case POSITION_FULL:
				client.gameState.QueueMessage(&FullPositionMessage{client.ID,message.ReadVector(),message.ReadVector(),message.ReadVector(),message.ReadVector()})
			case HIT:
				client.gameState.QueueMessage(&HitMessage{client.ID,message.ReadInt16(),message.ReadByte(),message.ReadInt16(),message.ReadVector()})
			case OPEN:
			case LOOT_ITEM:
				client.gameState.QueueMessage(&LootItemMessage{client.ID,message.ReadInt16(),message.ReadInt16(),message.ReadInt16()})
			case PICKUP_ITEM:
				client.gameState.QueueMessage(&PickupItemMessage{client.ID,message.ReadInt16()})
			case MOVE_ITEM:
				client.gameState.QueueMessage(&MoveItemMessage{client.ID,message.ReadInt16(),message.ReadInt16()})
			case USE_ITEM:
				client.gameState.QueueMessage(&UseItemMessage{client.ID,message.ReadInt16()})
			case NAME:
				client.gameState.QueueMessage(&NameMessage{client.ID,message.ReadString()})
			default:
				log.Printf("Bad message code: %X \n",messageValue)
			}
		}
	}

}

func (client *ClientConnection)Send(bytes []byte){
	binary.Write(client.conn,binary.LittleEndian,len(bytes))
	client.conn.Write(bytes)
}

