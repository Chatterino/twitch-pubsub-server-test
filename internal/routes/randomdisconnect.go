package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func RandomDisconnect(c *websocket.Conn, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), 1*time.Second)

	for {
		doBreak := defaultHandler(ctx, c, r)
		if doBreak {
			break
		}
	}
}
