package routes

type Message struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

var (
	pongPayload = []byte(`{"type":"PONG"}`)
)
