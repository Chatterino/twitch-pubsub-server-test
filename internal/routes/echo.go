package routes

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func Echo(c *websocket.Conn, r *http.Request) {
	for {
		ty, data, err := c.Read(r.Context())
		if err != nil {
			log.Printf("Failed to read from %v: %v", r.RemoteAddr, err)
			break
		}
		if string(data) == "/CLOSE" {
			err = c.Close(websocket.StatusNormalClosure, "Close command")
			if err != nil {
				log.Printf("Failed to gracefully close connection %v: %v", r.RemoteAddr, err)
			}
			break
		}
		err = c.Write(r.Context(), ty, data)
		if err != nil {
			log.Printf("Failed to write to %v: %v", r.RemoteAddr, err)
			break
		}
	}
}
