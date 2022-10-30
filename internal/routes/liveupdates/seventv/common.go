package seventv

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func ConvertMessage[D MessageType](c *Message[json.RawMessage]) (Message[D], error) {
	var d D
	err := json.Unmarshal(c.Data, &d)
	c2 := Message[D]{
		Op:        c.Op,
		Timestamp: c.Timestamp,
		Data:      d,
		Sequence:  c.Sequence,
	}

	return c2, err
}

func DefaultRead(ctx context.Context, c *websocket.Conn, r *http.Request) (*Message[json.RawMessage], bool) {
	_, data, err := c.Read(ctx)
	if err != nil {
		var websocketError websocket.CloseError
		if errors.As(err, &websocketError) {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				// Clean close
				return nil, false
			}
			log.Println("Unhandled close reason:", websocketError.Reason)
		}
		log.Println("failed to read message", r.RemoteAddr, err)
		return nil, false
	}

	var msg Message[json.RawMessage]
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println("Failed to unmarshal data", err)
		return nil, false
	}
	return &msg, true
}
