package liveupdates

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

func BasicSubUnsub(c *websocket.Conn, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	subs := make(map[string]int32)

	for {
		_, data, err := c.Read(ctx)
		if err != nil {
			var websocketError websocket.CloseError
			if errors.As(err, &websocketError) {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					// Clean close
					break
				}
				log.Println("Unhandled close reason:", websocketError.Reason)
			}
			log.Println("failed to read message", r.RemoteAddr, err)
			break
		}

		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("Failed to unmarshal data", err)
			break
		}

		switch msg.Op {
		case "sub":
			if _, ok := subs[msg.Condition]; !ok {
				subs[msg.Condition] = msg.Type
				c.Write(ctx, websocket.MessageText, []byte(fmt.Sprintf("ack-sub-%d-%s", msg.Type, msg.Condition)))
			}
		case "unsub":
			if v, ok := subs[msg.Condition]; ok && msg.Type == v {
				delete(subs, msg.Condition)
				c.Write(ctx, websocket.MessageText, []byte(fmt.Sprintf("ack-unsub-%d-%s", msg.Type, msg.Condition)))
			}
		default:
			log.Println("Unhandled message:", msg)
		}
	}
}
