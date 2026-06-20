package views

import (
	"time"
)

// RoomResponse representation
type RoomResponse struct {
	ID        string    `json:"room_id"`
	VideoURL  string    `json:"video_url"`
	StartedAt time.Time `json:"started_at"`
}
