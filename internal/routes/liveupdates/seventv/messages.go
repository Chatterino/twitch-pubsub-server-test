package seventv

import "encoding/json"

// MessageType - See definitions in
// https://github.com/SevenTV/API/blob/a400fe813ff9825fa79cd9d19cf2f0a3e8d42b32/data/events
type MessageType interface {
	json.RawMessage | HelloPayload | HeartbeatPayload |
		SubscribePayload | UnsubscribePayload
}

type Message[D MessageType] struct {
	Op        Opcode `json:"op"`
	Timestamp int64  `json:"t"`
	Data      D      `json:"d"`
	Sequence  uint64 `json:"s,omitempty"`
}

type Opcode uint8

const (
	OpcodeDispatch    Opcode = 0  // R - Server dispatches data to the client
	OpcodeHello       Opcode = 1  // R - Server greets the client
	OpcodeHeartbeat   Opcode = 2  // R - Keep the connection alive
	OpcodeReconnect   Opcode = 4  // R - Server demands that the client reconnects
	OpcodeSubscribe   Opcode = 35 // S - Subscribe to an event
	OpcodeUnsubscribe Opcode = 36 // S - Unsubscribe from an event
)

type eventType string

const (
	EventTypeUpdateEmoteSet eventType = "emote_set.update"
	EventTypeUpdateUser     eventType = "user.update"
)

type HelloPayload struct {
	HeartbeatInterval int64  `json:"heartbeat_interval"`
	SessionID         string `json:"session_id"`
}

type AckPayload struct {
	RequestID string         `json:"request_id"`
	Data      map[string]any `json:"data"`
}

type HeartbeatPayload struct {
	Count int64 `json:"count"`
}

type SubscribePayload struct {
	Type      eventType         `json:"type"`
	Condition map[string]string `json:"condition"`
}

type UnsubscribePayload struct {
	Type      eventType         `json:"type"`
	Condition map[string]string `json:"condition"`
}
