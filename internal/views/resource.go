package views

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
	Extra   any    `json:"extra,omitempty"`
}
