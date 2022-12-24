package bttv

import "encoding/json"

type MessageType interface {
	json.RawMessage | JoinPayload | PartPayload | BroadcastMePayload
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

type BroadcastMePayload struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
}
