package seventv

import (
	"context"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func NoHeartbeat(c *websocket.Conn, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	for {
		if _, ok := DefaultRead(ctx, c, r); !ok {
			break
		}
	}
}
