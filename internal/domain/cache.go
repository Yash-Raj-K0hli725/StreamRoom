package domain

import (
	"sync"
)

// Global In-Memory Database for demonstration
var (
	RoomsMap = make(map[string]*Room)
	RoomsMu  sync.RWMutex
)

func DeleteRoom(id string) {
	delete(RoomsMap, id)
}

func GetRoom(id string) *Room {
	return RoomsMap[id]
}
