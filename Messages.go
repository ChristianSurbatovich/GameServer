package main

type Message interface {
	GetType() byte
}

type PositionMessage struct{
	ID int16
	position vector
}

type VelocityMessage struct{
	ID int16
	velocity vector
}
