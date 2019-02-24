package main

import (
	"bytes"
	"encoding/binary"
)

func newMessage(length int16, positions int16, actions int16, buffer *bytes.Buffer){
	buffer.Reset()
	binary.Write(buffer,binary.LittleEndian,length)
	binary.Write(buffer,binary.LittleEndian,positions)
	binary.Write(buffer,binary.LittleEndian,actions)



}

func setHeader(length int16, positions int16, actions int16, buffer *bytes.Buffer){
	buffer.Reset()
	binary.Write(buffer,binary.LittleEndian,length)
	binary.Write(buffer,binary.LittleEndian,positions)
	binary.Write(buffer,binary.LittleEndian,actions)
}

func createAction(command byte, value int16, targetState *ShipData, transform *Transform) []byte{
		buffer := new(bytes.Buffer)
		switch command{
		case HIT:
			break
		case HEALTH:
			break
		case SINK:
			break
		case FIRE:
			break
		case OPEN:
			break
		case NAME:
			break
		case SPAWN:
			binary.Write(buffer,binary.LittleEndian,SPAWN)
			binary.Write(buffer,binary.LittleEndian,transform.position.x)
			binary.Write(buffer,binary.LittleEndian,transform.position.y)
			binary.Write(buffer,binary.LittleEndian,transform.position.z)
			binary.Write(buffer,binary.LittleEndian,transform.rotation.x)
			binary.Write(buffer,binary.LittleEndian,transform.rotation.y)
			binary.Write(buffer,binary.LittleEndian,transform.rotation.z)
			binary.Write(buffer,binary.LittleEndian,targetState.ID)
			binary.Write(buffer,binary.LittleEndian,value)
			binary.Write(buffer,binary.LittleEndian,targetState.doorsOpen)
			binary.Write(buffer,binary.LittleEndian,clientStats[targetState.ID].baseStats[CURRENT_HEALTH])
			binary.Write(buffer,binary.LittleEndian,clientStats[targetState.ID].totalStats[MAX_HEALTH])
			binary.Write(buffer,binary.LittleEndian,int16(len(targetState.name)))
			buffer.Write([]byte(targetState.name))
			break
		case REGISTER:
			break
		case RESPAWN:
			break
		case REMOVE:
			break
		case DESPAWN:
			break
		case CHAT:
			break
		case FEED:
			break
		case EXPLODE:
			break
		case POSITION:
			break
		case TEST:
			break
		}
		return buffer.Bytes()
}

func addMessage(command byte, value int16, targetState *ShipData, transform *Transform, buffer *bytes.Buffer){
	switch command{
	case REGISTER:
		buffer.WriteByte(REGISTER)
		binary.Write(buffer,binary.LittleEndian,targetState.ID)
		binary.Write(buffer,binary.LittleEndian,serverTickRate)
		break
	case SPAWN:
		binary.Write(buffer,binary.LittleEndian,SPAWN)
		binary.Write(buffer,binary.LittleEndian,transform.position.x)
		binary.Write(buffer,binary.LittleEndian,transform.position.y)
		binary.Write(buffer,binary.LittleEndian,transform.position.z)
		binary.Write(buffer,binary.LittleEndian,transform.rotation.x)
		binary.Write(buffer,binary.LittleEndian,transform.rotation.y)
		binary.Write(buffer,binary.LittleEndian,transform.rotation.z)
		binary.Write(buffer,binary.LittleEndian,targetState.ID)
		binary.Write(buffer,binary.LittleEndian,value)
		binary.Write(buffer,binary.LittleEndian,targetState.doorsOpen)
		binary.Write(buffer,binary.LittleEndian,clientStats[targetState.ID].baseStats[CURRENT_HEALTH])
		binary.Write(buffer,binary.LittleEndian,clientStats[targetState.ID].totalStats[MAX_HEALTH])
		binary.Write(buffer,binary.LittleEndian,int16(len(targetState.name)))
		buffer.Write([]byte(targetState.name))
		break
	}
}