package main

const(
	HIT byte = 0x01
	HEALTH byte = 0x02
	SINK byte = 0x03
	FIRE byte = 0x04
	OPEN byte = 0x05
	NAME byte = 0x06
	SPAWN byte = 0x07
	REGISTER byte = 0x08
	RESPAWN byte = 0x09
	REMOVE byte = 0x0A
	DESPAWN byte = 0x0B
	CHAT byte = 0x0C
	FEED byte = 0x0D
	EXPLODE byte = 0x0E
	POSITION byte = 0x0F
	STAT byte = 0x10
	EQUIP_ITEM byte = 0x13
	LOOT_AREA byte = 0x15
	LOOT_ITEM byte = 0x16
	ADD_LOOT_ITEM byte = 0x17
	PICKUP_ITEM byte = 0x18
	UNEQUIP_ITEM byte = 0x19
	SYNC byte = 0x1A
	DESTRUCTION_STATE_RESET byte = 0x1B
	DESTRUCTION_STATE byte = 0x1C
	ADD_ITEM byte = 0x1D
	USE_ITEM byte = 0x1E
	ADD_STATIC_OBJECT byte = 0x1F
	ADD_DYNAMIC_OBJECT byte = 0x20
	ADD_STATIC_ACTOR byte = 0x21
	ADD_DYNAMIC_ACTOR byte = 0x22
	MOVE_ITEM byte = 0x23
	REMOVE_ITEM byte = 0x24
	REMOVE_LOOT_ITEM byte = 0x25
	POSITION_FULL byte = 0x26
	VELOCITY byte = 0x27
	VELOCITY_FULL byte = 0x28
	AGENT_SYNC byte  = 0x29
	ADD_AGENT byte = 0x2A
	FORCE_ITEM_USE byte = 0x2B
	CONNECT byte = 0x2C
	RECONNECT byte = 0x2D
	TEST byte = 0xFF


	HIT_LENGTH int16 = 18
	HEALTH_LENGTH int16 = 9
	SINK_LENGTH int16 = 29
	REGISTER_LENGTH int16 = 9
	SPAWN_LENGTH int16 = 40
)
