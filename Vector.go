package main

import "math"

type Vector struct{
	x float32
	y float32
	z float32
}

func (v1 *Vector)Sub(v2 Vector) (result Vector){
	result.x = v1.x - v2.x
	result.y = v1.y - v2.y
	result.z = v1.z - v2.z
	return
}

func (v1 *Vector)DistanceSquared(v2 Vector) float32{
	difference := v1.Sub(v2)
	return difference.x * difference.x + difference.y * difference.y + difference.z * difference.z
}

func (v1 *Vector)Distance(v2 Vector)float32{
	difference := v1.Sub(v2)
	distance2 := difference.x * difference.x + difference.y * difference.y + difference.z * difference.z
	return float32(math.Sqrt(float64(distance2)))
}
