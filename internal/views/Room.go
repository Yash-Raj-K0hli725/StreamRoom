package views

import "time"

// Room representation
type Room struct {
	ID        string    `json:"room_id"`
	VideoURL  string    `json:"video_url"`
	StartedAt time.Time `json:"started_at"`
}
