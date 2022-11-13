package bttv

import "encoding/json"

type MessageType interface {
	json.RawMessage | JoinPayload | PartPayload
}

type Message[D MessageType] struct {
	Name string `json:"name"`
	Data D      `json:"data"`
}

type JoinPayload struct {
	Name string `json:"name"`
}

type PartPayload struct {
	Name string `json:"name"`
}
