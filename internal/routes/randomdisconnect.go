package routes

import (
	"context"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

func RandomDisconnect(c *websocket.Conn, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 1500*time.Millisecond)

	for {
		doBreak := defaultHandler(ctx, c, r)
		if doBreak {
			break
		}
	}
}
