package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

func Echo(c *websocket.Conn, r *http.Request) {
	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	for {
		err := echoMessage(r.Context(), c, l)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			log.Printf("failed to echo with %v: %v", r.RemoteAddr, err)
			return
		}
	}
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echoMessage(ctx context.Context, c *websocket.Conn, l *rate.Limiter) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err := l.Wait(ctx)
	if err != nil {
		return err
	}

	typ, r, err := c.Read(ctx)
	if err != nil {
		return err
	}

	log.Println("/echo Got message:", string(r))

	err = c.Write(ctx, typ, r)
	return err
}
