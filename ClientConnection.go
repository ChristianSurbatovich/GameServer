package main

import (
	"net"
	"sync"
	"io"
	"encoding/binary"
	"bytes"
	"bufio"
	"log"
)

type ClientConnection struct{
	conn net.Conn
	lock *sync.RWMutex
	reader *bufio.Reader
	lengthBuffer []byte
	messageBuffer *bytes.Buffer
	messageLength int16
}

func NewClientConnection(newConnection net.Conn) *ClientConnection{
	return &ClientConnection{
		newConnection,
		&sync.RWMutex{},
		bufio.NewReader(newConnection),
		make([]byte,PREFIX_SIZE),
		new(bytes.Buffer),
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
		message := make([]byte,client.messageLength)
		n, err = client.reader.Read(message)
		if err != nil{
			log.Println(err)
			if err == io.EOF{
				continue
			}else{
				break
			}
		}
		if len(message) > 1 {
			switch message[0] {
			case SYNC:

			}
		}
	}

}