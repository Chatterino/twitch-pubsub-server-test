package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func Default(c *websocket.Conn, r *http.Request) {
	for {
		_, data, err := c.Read(r.Context())
		if err != nil {
			var websocketError *websocket.CloseError
			if errors.As(err, &websocketError) {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {

				}
				log.Println("Close reason:", websocketError.Reason)
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
			if err := c.Write(r.Context(), websocket.MessageText, pongPayload); err != nil {
				log.Println("Error writing response", err)
				return
			}
		}

		log.Println("/: Got message:", msg)
	}
}
