package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"sync"
)












type Player struct{
	playerData *PlayerData
	clientData ClientData
	messageLength int16
	playerID int16
	abilities map[int16]ability
	localIDs int16
	lock *sync.Mutex

	rnd *rand.Rand

}

func NewPlayer(id int16, conn net.Conn)*Player{
	var player Player
	player.playerID = id
	player.playerData = NewPlayerData()
	player.clientData = NewClientData(conn)

	player.abilities = make(map[int16]ability)
	player.messageBuffer = new(bytes.Buffer)
	player.messageBuffer.Grow(256)
	player.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	player.reader = bufio.NewReader(conn)
	return &player
}

func (player *Player) mainLoop(){
	for{

		if len(message) > 1{
			switch message[0]{
			case SYNC:
				player.messageBuffer.Reset()
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(9))
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(0))
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(1))
				player.messageBuffer.WriteByte(SYNC)
				binary.Write(player.messageBuffer,binary.LittleEndian,float32(time.Now().Sub(serverTime).Seconds()))
				player.clientData.lock.Lock()
				player.clientData.conn.Write(player.messageBuffer.Bytes())
				player.clientData.lock.Unlock()

			case POSITION:
				if player.playerData.state.state == 1{
					continue
				}
				messageReader := bytes.NewReader(message[3:])
				player.playerData.lock.RLock()
				transform := player.playerData.transform
				binary.Read(messageReader,binary.LittleEndian,&transform.position.x)
				binary.Read(messageReader,binary.LittleEndian,&transform.position.y)
				binary.Read(messageReader,binary.LittleEndian,&transform.position.z)
				binary.Read(messageReader,binary.LittleEndian,&transform.rotation.x)
				binary.Read(messageReader,binary.LittleEndian,&transform.rotation.y)
				binary.Read(messageReader,binary.LittleEndian,&transform.rotation.z)
				binary.Read(messageReader,binary.LittleEndian,&transform.velocity.x)
				binary.Read(messageReader,binary.LittleEndian,&transform.velocity.y)
				binary.Read(messageReader,binary.LittleEndian,&transform.velocity.z)
				binary.Read(messageReader,binary.LittleEndian,&transform.locationTime)
				player.playerData.lock.RUnlock()
				// fire id numshots x1 y1 z1 x2 y2 z2 ...
			case FIRE:
				actions.Push(message)
			case OPEN:
				actions.Push(message)
				player.playerData.state.doorsOpen = !player.playerData.state.doorsOpen
			case HIT:
				var playerHit bool
				var id int16
				var weaponID int16
				messageReader := bytes.NewReader(message[1:])
				binary.Read(messageReader,binary.LittleEndian,&playerHit)
				binary.Read(messageReader,binary.LittleEndian,&id)
				binary.Read(messageReader,binary.LittleEndian,&weaponID)
				if playerHit {
					temp := players[id]
					if temp != nil {
						temp.playerData.stats.totalStats[CURRENT_HEALTH] -= player.playerData.stats.totalStats[weaponID]
						if temp.playerData.stats.totalStats[CURRENT_HEALTH] <= 0 && temp.playerData.state.alive {
							temp.playerData.state.alive = false
							action := new(bytes.Buffer)
							action.WriteByte(FEED)
							binary.Write(action, binary.LittleEndian, player.playerID)
							binary.Write(action, binary.LittleEndian, id)
							actions.Push(action.Bytes())
						}
					}
					if id == player.playerID && player.playerData.stats.totalStats[CURRENT_HEALTH] <= 0 {
						player.playerData.state.state = 1
					}
					action := new(bytes.Buffer)
					action.WriteByte(HEALTH)
					binary.Write(action,binary.LittleEndian,id)
					binary.Write(action,binary.LittleEndian,temp.playerData.stats.totalStats[CURRENT_HEALTH])
					actions.Push(action.Bytes())
				}
				actions.Push(message)
			case NAME:
				var id, nameLength int16
				messageReader := bytes.NewReader(message[1:])
				binary.Read(messageReader,binary.LittleEndian,&id)
				binary.Read(messageReader,binary.LittleEndian,&nameLength)
				player.playerData.state.name = string(message[5:5+nameLength])
				actions.Push(message)
			case SPAWN:
				if player.playerData.state.canSpawn{
					sL := <-nextSpawn
					player.playerData.transform.position = Vector{spawnLocations[sL].x,spawnLocations[sL].y,spawnLocations[sL].z}
					player.playerData.transform.rotation = Vector{0,float32(player.rnd.Intn(360)),0}
					player.playerData.stats.totalStats[CURRENT_HEALTH] = player.playerData.stats.totalStats[MAX_HEALTH]
					player.playerData.state.doorsOpen = false
					player.playerData.state.alive = true
					player.playerData.state.state = ALIVE
					actions.Push(createAction(SPAWN,0,player.playerData.state,player.playerData.transform))
				}
			case CHAT:
				var chatLength, nameLength int16
				nameLength = int16(len(player.playerData.state.name))
				chatLength = int16(len(message) - 3)
				action := new(bytes.Buffer)
				action.WriteByte(CHAT)
				binary.Write(action,binary.LittleEndian,nameLength)
				binary.Write(action,binary.LittleEndian,chatLength)
				action.Write([]byte(player.playerData.state.name))
				action.Write(message[3:])
				actions.Push(action.Bytes())
			case ABILITY:
				var abilityID int16
				var location Vector
				reader := bytes.NewReader(message[1:])
				binary.Read(reader,binary.LittleEndian,&abilityID)
				binary.Read(reader,binary.LittleEndian,&location)
				abilities[abilityID].Use(clientID,location)
			case LOOT_AREA:
				var areaID int16
				binary.Read(bytes.NewReader(message[1:]),binary.LittleEndian,&areaID)
				if !player.playerData.items.lootAreas[areaID].looted{
					player.playerData.items.lootAreas[areaID].generateLoot(itemList)
				}
				player.messageBuffer.Reset()
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(len(player.playerData.items.lootAreas[areaID].lootList) * 4 + 7))
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(0))
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(1))
				player.messageBuffer.WriteByte(ADD_LOOT_ITEM)
				binary.Write(player.messageBuffer,binary.LittleEndian,int16(len(player.playerData.items.lootAreas[areaID].lootList)))
				for localID, item := range player.playerData.items.lootAreas[areaID].lootList{
					binary.Write(player.messageBuffer,binary.LittleEndian,item.id())
					binary.Write(player.messageBuffer,binary.LittleEndian,localID)
				}
				player.clientData.conn.Write(player.messageBuffer.Bytes())
			case PICKUP_ITEM:
				var localID int16
				var areaID int16
				binary.Read(bytes.NewReader(message[1:]),binary.LittleEndian,&localID)
				binary.Read(bytes.NewReader(message[3:]),binary.LittleEndian,&areaID)
				itemLocation, status := player.playerData.items.PickupItem(areaID,localID)
				if status == PICKUP_SUCCESS{
					player.messageBuffer.Reset()
					binary.Write(player.messageBuffer,binary.LittleEndian,int16(9))
					binary.Write(player.messageBuffer,binary.LittleEndian,int16(0))
					binary.Write(player.messageBuffer,binary.LittleEndian,int16(1))
					player.messageBuffer.WriteByte(LOOT_ITEM)
					binary.Write(player.messageBuffer,binary.LittleEndian,localID)
					binary.Write(player.messageBuffer,binary.LittleEndian,areaID)
					binary.Write(player.messageBuffer,binary.LittleEndian,itemLocation)
					player.clientData.conn.Write(player.messageBuffer.Bytes())
				}
			case EQUIP_ITEM:
				var source, destination int16
				binary.Read(bytes.NewReader(message[1:]),binary.LittleEndian,&source)
				binary.Read(bytes.NewReader(message[3:]),binary.LittleEndian,&destination)
				client.player.items.swap(source,destination)
				if item, exists := inventoryItems[localID]; exists{
					equippedItems[localID] = item
					delete(inventoryItems,localID)
					item.onEquip(clientID)
				}
				item := itemMap[localID]
				item.itemID = localID
			case UNEQUIP_ITEM:
				var localID int16
				binary.Read(bytes.NewReader(message[1:]),binary.LittleEndian,&localID)
				if item, exists := equippedItems[localID]; exists{
					inventoryItems[localID] = item
					delete(equippedItems,localID)
					item.onUnequip(clientID)
				}
			case DESTRUCTION_STATE_RESET:
				shipStates[clientID].destructionState = nil
				actions.Push(message)
			case DESTRUCTION_STATE:
				shipStates[clientID].destructionState = append(shipStates[clientID].destructionState,message[3:]...)
			case ADD_ABILITY:
			case EQUIP_ABILITY:
			case REMOVE_ABILITY:
			case UNEQUIP_ABILITY:

			}
		}
	}
}

func (player *Player) initialize(conn net.Conn, id int16){

	fmt.Println("Received a connection from: " + conn.RemoteAddr().String())
	log.Println("Received a connection from: " + conn.RemoteAddr().String())


	elem, ok := clients[-1]
	if ok{
		elem.client = conn
		newMessage(REGISTER_LENGTH,0,1,client.messageBuffer)
		addMessage(REGISTER,0,shipStates[client.clientData.clientID],nil,client.messageBuffer)
		conn.Write(client.messageBuffer.Bytes())
	}else{
		client.clientData.clientID = id
		sL := <-nextSpawn
		client.player.transform = Transform{Vector{spawnLocations[sL].x,spawnLocations[sL].y,spawnLocations[sL].z}, Vector{0,float32(client.rnd.Intn(360)),0}, Vector{0,0,0},client.clientData.clientID,0}

		client.player.state = ShipData{id,false,nil,"",ALIVE,time.Now(),true,true}

		newMessage(REGISTER_LENGTH,0,1,client.messageBuffer)
		addMessage(REGISTER,0, &client.player.state,nil,client.messageBuffer)
		conn.Write(client.messageBuffer.Bytes())
	}
	actions.Push(createAction(SPAWN,0,&client.player.state,&client.player.transform))
	client.messageBuffer.Reset()
	newMessage(0,0,int16((len(clients) - 1) * 2),client.messageBuffer)
	stateLock.Lock()
	for _, pt := range clients{
		if pt.agentID == id || shipStates[pt.agentID].alive == false{
			continue
		}
		addMessage(SPAWN,0,shipStates[pt.agentID],pt.transform,client.messageBuffer)
		client.messageBuffer.WriteByte(DESTRUCTION_STATE)
		binary.Write(client.messageBuffer,binary.LittleEndian,pt.agentID)
		binary.Write(client.messageBuffer,binary.LittleEndian,int16(len(shipStates[pt.agentID].destructionState)))
		client.messageBuffer.Write(shipStates[pt.agentID].destructionState)
	}
	stateLock.Unlock()
	byteMessage := client.messageBuffer.Bytes()
	client.messageBuffer.Reset()
	binary.Write(client.messageBuffer,binary.LittleEndian,int16(len(byteMessage) - 2))
	conn.Write(byteMessage)
}

func (client *client) close(){

}

func handleClient(conn net.Conn){





	log.Printf("Player %s disconnected",shipStates[clientID].name)
	stateLock.Lock()
	delete(shipStates, clientID)
	delete(playerStates,clientID)
	delete(clients,clientID)
	action := new(bytes.Buffer)
	action.WriteByte(REMOVE)
	binary.Write(action,binary.LittleEndian,clientID)
	actions.Push(action.Bytes())
	stateLock.Unlock()
}