package views

type ActionRequest struct {
	Action     string `json:"action"`      // "PLAY", "PAUSE", "SEEK"
	PositionMs int64  `json:"position_ms"` // Used for SEEK or syncing state
}
