package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func handleListen(ctx context.Context, c *websocket.Conn, r *http.Request, msg Message) {
	response := Response{
		Type:  "RESPONSE",
		Nonce: msg.Nonce,
		Error: "",
	}

	write(ctx, c, r, &response)
}

func handleListenBadAuth(ctx context.Context, c *websocket.Conn, r *http.Request, msg Message) {
	response := Response{
		Type:  "RESPONSE",
		Nonce: msg.Nonce,
		Error: "ERR_BADAUTH",
	}

	write(ctx, c, r, &response)
}

func defaultHandler(ctx context.Context, c *websocket.Conn, r *http.Request) bool {
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
		log.Println("defaultHandler failed to read message", r.RemoteAddr, err)
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
		handleListen(ctx, c, r, msg)
	default:
		log.Println("Unhandled message:", msg)
	}

	return false
}

func handlePing(ctx context.Context, c *websocket.Conn, r *http.Request) {
	err := c.Write(ctx, websocket.MessageText, pongPayload)
	if err != nil {
		panic(err)
	}
}

func write(ctx context.Context, c *websocket.Conn, r *http.Request, msg interface{}) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = c.Write(ctx, websocket.MessageText, bytes)
	if err != nil {
		panic(err)
	}
}
