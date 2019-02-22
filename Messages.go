package main

type Message interface {
	GetType() byte
}


type SyncMessage struct{
	serverTime float32
}

func (message *SyncMessage)GetType()byte{
	return SYNC
}


type PositionMessage struct{
	actorID int16
	position vector
	rotation vector
}

func (message *PositionMessage)GetType()byte{
	return POSITION
}

type VelocityMessage struct{
	actorID int16
	velocity vector
	angularVelocity vector
}

func (message *VelocityMessage)GetType()byte{
	return VELOCITY
}

type FullPositionMessage struct{
	actorID int16
	position vector
	rotation vector
	velocity vector
	angularVelocity vector
}

func(message *FullPositionMessage)GetType()byte{
	return POSITION_FULL
}

type HitMessage struct{
	reporter int16
	eventID int16
	targetType byte
	targetID int16
	offset vector
}

func(message *HitMessage)GetType()byte{
	return HIT
}

type UseItemMessage struct{
	actorID int16
	itemSlotID int16
}

func(message *UseItemMessage)GetType()byte{
	return USE_ITEM
}

type PickupItemMessage struct{
	actorID int16
	itemPickupID int16
}

func(message *PickupItemMessage)GetType()byte{
	return PICKUP_ITEM
}

type LootItemMessage struct{
	actorID int16
	areaID int16
	itemID int16
	itemSlotID int16
}

func(message *LootItemMessage)GetType()byte{
	return LOOT_ITEM
}

type MoveItemMessage struct{
	actorID int16
	itemSlot1 int16
	itemSlot2 int16
}

func(message *MoveItemMessage)GetType()byte{
	return MOVE_ITEM
}

type NameMessage struct{
	actorID int16
	name string
}

func (message *NameMessage)GetType()byte{
	return NAME
}