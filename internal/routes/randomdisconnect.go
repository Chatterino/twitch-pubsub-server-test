package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

func RandomDisconnect(c *websocket.Conn, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 1500*time.Millisecond)
	for {
		_, data, err := c.Read(ctx)
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
			if err := c.Write(ctx, websocket.MessageText, pongPayload); err != nil {
				log.Println("Error writing response", err)
				return
			}
		}

		log.Println("/: Got message:", msg)
	}
}
