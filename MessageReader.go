package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

type MessageReader struct{
	reader *bytes.Reader
}


func NewMessageReader(byteslice []byte) *MessageReader {
	return &MessageReader{reader:bytes.NewReader(byteslice)}
}

func (buffer *MessageReader)UnreadData()bool{
	return buffer.reader.Len() > 0
}

func (buffer *MessageReader)ReadByte() byte{
	value, err :=buffer.reader.ReadByte()
	if err != nil {
		log.Println(err)
	}
	return value
}

func (buffer *MessageReader)ReadInt16() int16{
	var value int16
	err := binary.Read(buffer.reader,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (buffer *MessageReader)ReadInt() int{
	var value int
	err := binary.Read(buffer.reader,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (buffer *MessageReader)ReadFloat32() float32{
	var value float32
	err := binary.Read(buffer.reader,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (buffer *MessageReader)ReadFloat64() float64{
	var value float64
	err := binary.Read(buffer.reader,binary.LittleEndian,&value)
	if err != nil {
		log.Println(err)
	}
	return value
}

func (buffer *MessageReader)ReadVector() vector{
	return vector{buffer.ReadFloat32(),buffer.ReadFloat32(),buffer.ReadFloat32()}
}

func (buffer *MessageReader)ReadString() string{
	length := buffer.ReadInt16()
	stringBytes := make([]byte,length)
	buffer.reader.Read(stringBytes)
	return string(stringBytes)
}