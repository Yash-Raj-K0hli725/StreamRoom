package views

type SyncMessage struct {
	Event     string `json:"event"`
	CurrentMS int    `json:"ms"`
}
