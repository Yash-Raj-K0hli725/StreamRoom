package domain

import (
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

type Room struct {
	ID      string
	Clients map[*Client]bool
	Mu      sync.Mutex
	Ctx     context.Context
	Cancel  context.CancelFunc

	// Shared Video State
	IsPlaying         bool
	CurrentPositionMs int64 // The anchored playback time
	LastUpdated       time.Time
}

// GetLivePosition calculates the real-time position safely
func (r *Room) GetLivePosition() int64 {
	if !r.IsPlaying {
		return r.CurrentPositionMs
	}
	// If playing, position is anchor position + time elapsed since last anchor change
	return r.CurrentPositionMs + time.Since(r.LastUpdated).Milliseconds()
}
