package main


type TaggedMessage struct{
	position Vector
	neverIgnore bool
	ignoreRange float32
	bytes []byte
}

func (tag *TaggedMessage)CheckTag(v Vector) bool{
	return tag.neverIgnore || tag.position.DistanceSquared(v) < tag.ignoreRange
}