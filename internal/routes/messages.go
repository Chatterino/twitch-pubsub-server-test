package routes

type Message struct {
	Type  string                 `json:"type"`
	Nonce string                 `json:"nonce"`
	Data  map[string]interface{} `json:"data"`
}

type Response struct {
	Type  string `json:"type"`
	Nonce string `json:"nonce"`
	Error string `json:"error"`
}

var (
	pongPayload = []byte(`{"type":"PONG"}`)
)
