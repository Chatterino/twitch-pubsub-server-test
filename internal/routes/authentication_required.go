package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func authenticatedHandler(ctx context.Context, c *websocket.Conn, r *http.Request) bool {
	const correctAuthToken = "xD"

	_, data, err := c.Read(ctx)
	if err != nil {
		var websocketError websocket.CloseError
		if errors.As(err, &websocketError) {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				// Clean close
				return true
			}
			log.Println("Unhandled close reason:", websocketError.Reason)
		}
		log.Println("authenticatedHandler failed to read message", r.RemoteAddr, err)
		return true
	}

	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println("Failed to unmarshal data", err)
		return true
	}

	switch msg.Type {
	case "PING":
		handlePing(ctx, c, r)
	case "LISTEN":
		if authTokenV, ok := msg.Data["auth_token"]; !ok {
			fmt.Println("Missing auth_token field")
			// Missing auth_token field
			handleListenBadAuth(ctx, c, r, msg)
		} else {
			if authToken, ok := authTokenV.(string); !ok {
				fmt.Println("auth_token wrong type")
				// auth_token wrong type
				handleListenBadAuth(ctx, c, r, msg)
			} else {
				if authToken == correctAuthToken {
					handleListen(ctx, c, r, msg)
				} else {
					fmt.Println("auth_token wrong value, should be", correctAuthToken)
					handleListenBadAuth(ctx, c, r, msg)
				}
			}
		}
	default:
		log.Println("Unhandled message:", msg)
	}

	return false
}

func AuthenticationRequired(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()

	for {
		doBreak := authenticatedHandler(ctx, c, r)
		if doBreak {
			break
		}
	}
}
