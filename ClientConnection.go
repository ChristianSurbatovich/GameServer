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
		message := &RWMessage{bytes.NewBuffer(messageBytes)}
		switch message.CheckByte(){
		case SYNC:
			syncMessage := &RWMessage{new(bytes.Buffer)}
			syncMessage.WriteByte(SYNC)
			syncMessage.Write(float32(time.Now().Sub(serverStartTime).Seconds()))
			client.Send(syncMessage.Bytes())
		default:
			gameState.QueueMessage(message)
		}
	}
}

func (client *ClientConnection)Send(bytes []byte){
	binary.Write(client.conn,binary.LittleEndian,len(bytes))
	client.conn.Write(bytes)
}

