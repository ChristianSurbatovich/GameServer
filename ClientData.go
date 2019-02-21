package main

import (
	"net"
	"sync"
)

type ClientData struct{
	ID int16
}

func NewClientData(newConnection net.Conn) ClientData{
	return ClientData{conn:newConnection,lock:&sync.RWMutex{}}
}