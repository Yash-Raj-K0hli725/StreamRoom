package service

import (
	"StreamRoom/internal/cache"
	"StreamRoom/internal/views"
	"fmt"
	"time"
)

// RoomService handles generation and storage of synced rooms
type RoomService struct{}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (s *RoomService) CreateRoom(videoURL string) *views.Room {
	cache.RoomsMu.Lock()
	defer cache.RoomsMu.Unlock()

	// Generate a short, unique alphanumeric room code
	roomID := fmt.Sprintf("ROOM-%d", time.Now().UnixNano()%100000)

	newRoom := &views.Room{
		ID:        roomID,
		VideoURL:  videoURL,
		StartedAt: time.Now(), // The live broadcast ticker clock begins ticking NOW
	}

	cache.RoomsMap[roomID] = newRoom

	// Note: In your final system, you would trigger the background synchronization loop
	// (like the ticker we designed earlier) right here using: go runRoomSyncLoop(newRoom)

	return newRoom
}
