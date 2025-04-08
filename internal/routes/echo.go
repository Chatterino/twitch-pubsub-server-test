package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/coder/websocket"
)

func Echo(c *websocket.Conn, r *http.Request) {
	for {
		ty, data, err := c.Read(r.Context())
		if err != nil {
			log.Printf("Failed to read from %v: %v", r.RemoteAddr, err)
			break
		}
		strData := string(data)

		if strData == "/CLOSE" {
			err = c.Close(websocket.StatusNormalClosure, "Close command")
			if err != nil {
				log.Printf("Failed to gracefully close connection %v: %v", r.RemoteAddr, err)
			}
			break
		}
		if strings.HasPrefix(strData, "/HEADER ") {
			data = []byte(r.Header.Get(strings.TrimPrefix(strData, "/HEADER ")))
		} else if strData == "/URL" {
			data = []byte(r.URL.String())
		}

		err = c.Write(r.Context(), ty, data)
		if err != nil {
			log.Printf("Failed to write to %v: %v", r.RemoteAddr, err)
			break
		}
	}
}
