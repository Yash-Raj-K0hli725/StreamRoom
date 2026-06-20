package views

type SyncResponse struct {
	IsPlaying bool  `json:"is_playing"`
	CurrentMS int64 `json:"ms"`
}
