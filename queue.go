package main

import "sync"

type Queue struct{
	front *Node
	back *Node
	dataLock *sync.Mutex
	size int
}

type Node struct{
	value interface{}
	next *Node
}

func NewQueue() *Queue{
	return &Queue{dataLock:&sync.Mutex{}}
}

func (queue *Queue) Push(value interface{}){
	queue.dataLock.Lock()
	defer queue.dataLock.Unlock()
	queue.size++
	if queue.back != nil{
		queue.back.next = &Node{value,nil}
		queue.back = queue.back.next
		return
	}
	queue.back = &Node{value:value}
	queue.front = queue.back
}

func (queue *Queue) Peek() (interface{},bool){
	queue.dataLock.Lock()
	defer queue.dataLock.Unlock()
	if queue.front != nil {
		return queue.front.value,true
	}
	return nil,false
}

func (queue *Queue) Pop() interface{}{
	queue.dataLock.Lock()
	defer queue.dataLock.Unlock()
	if queue.front == nil {
		return []byte{}
	}
	queue.size--
	v := queue.front
	queue.front = v.next
	if queue.front == nil{
		queue.back = nil
	}
	return v.value
}

func (queue *Queue) Size() int{
	queue.dataLock.Lock()
	defer queue.dataLock.Unlock()
	return queue.size
}