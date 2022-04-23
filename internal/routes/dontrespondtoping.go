package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func DontRespondToPing(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()

	for {
		_, data, err := c.Read(ctx)
		if err != nil {
			var websocketError websocket.CloseError
			if errors.As(err, &websocketError) {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					// Clean close
					return
				}
				log.Println("Unhandled close reason:", websocketError.Reason)
			}
			log.Printf("failed to read message %v: %v", r.RemoteAddr, err)
			return
		}

		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("Failed to unmarshal data", err)
			return
		}

		switch msg.Type {
		case "PING":
			// handlePing(ctx, c, r)
		case "LISTEN":
			handleListen(ctx, c, r, msg)
		default:
			log.Println("Unhandled message:", msg)
		}
	}
}
