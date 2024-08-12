package routes

import (
	"net/http"

	"github.com/coder/websocket"
)

func Default(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()

	for {
		doBreak := defaultHandler(ctx, c, r)
		if doBreak {
			break
		}
	}
}
