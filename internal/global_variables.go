package internal

import (
	"StreamRoom/internal/views"
	"sync"
)

// Global In-Memory Database for demonstration
var (
	RoomsMap = make(map[string]*views.Room)
	RoomsMu  sync.RWMutex
)
