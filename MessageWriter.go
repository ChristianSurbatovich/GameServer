package main

import (
	"bytes"
	"encoding/binary"
)

type MessageWriter struct{
	buffer *bytes.Buffer
}

func NewMessageWriter() *MessageWriter{
	return &MessageWriter{new(bytes.Buffer)}
}

func (message *MessageWriter)Bytes() []byte{
	return message.buffer.Bytes()
}

func (message *MessageWriter)Write(value interface{}){
	binary.Write(message.buffer,binary.LittleEndian,value)
}

func (message *MessageWriter)WriteByte(value byte){
	message.buffer.WriteByte(value)
}

func (message *MessageWriter)WriteVector(v vector){
	binary.Write(message.buffer,binary.LittleEndian,v.x)
	binary.Write(message.buffer,binary.LittleEndian,v.y)
	binary.Write(message.buffer,binary.LittleEndian,v.z)
}

func (message *MessageWriter)WriteString(s string){
	stringBytes := []byte(s)
	message.Write(int16(len(stringBytes)))
	for _,b := range stringBytes{
		message.WriteByte(b)
	}
}