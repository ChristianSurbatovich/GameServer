package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

type RWMessage struct{
	buffer *bytes.Buffer
}

func NewRWMessage(initialBuffer *bytes.Buffer) *RWMessage{
	return &RWMessage{buffer:initialBuffer}
}

func (message *RWMessage)Reset(){
	message.buffer.Reset()
}

func (message *RWMessage)UnreadData()bool{
	return message.buffer.Len() > 0
}

func (message *RWMessage)CheckByte() byte{
	value, err := message.buffer.ReadByte()
	if err != nil {
		log.Println(err)
		return value
	}
	message.buffer.UnreadByte()
	return value
}

func (message *RWMessage)ReadByte() byte{
	value, err :=message.buffer.ReadByte()
	if err != nil {
		log.Println(err)
	}
	return value
}

func (message *RWMessage)ReadInt16() int16{
	var value int16
	err := binary.Read(message.buffer,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (message *RWMessage)ReadInt() int{
	var value int
	err := binary.Read(message.buffer,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (message *RWMessage)ReadFloat32() float32{
	var value float32
	err := binary.Read(message.buffer,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (message *RWMessage)ReadFloat64() float64{
	var value float64
	err := binary.Read(message.buffer,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (message *RWMessage)ReadVector() Vector {
	return Vector{message.ReadFloat32(),message.ReadFloat32(),message.ReadFloat32()}
}

func (message *RWMessage)ReadString() string{
	length := message.ReadInt16()
	stringBytes := make([]byte,length)
	message.buffer.Read(stringBytes)
	return string(stringBytes)
}

func (message *RWMessage)Bytes() []byte{
	return message.buffer.Bytes()
}

func (message *RWMessage)Write(value interface{}){
	binary.Write(message.buffer,binary.LittleEndian,value)
}

func (message *RWMessage)WriteByte(value byte){
	message.buffer.WriteByte(value)
}

func (message *RWMessage)WriteVector(v Vector){
	binary.Write(message.buffer,binary.LittleEndian,v.x)
	binary.Write(message.buffer,binary.LittleEndian,v.y)
	binary.Write(message.buffer,binary.LittleEndian,v.z)
}

func (message *RWMessage)WriteString(s string){
	stringBytes := []byte(s)
	message.Write(int16(len(stringBytes)))
	for _,b := range stringBytes{
		message.WriteByte(b)
	}
}

func (message *RWMessage)String()string{
	return message.buffer.String()
}