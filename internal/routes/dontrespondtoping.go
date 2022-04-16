package routes

import (
	"errors"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func DontRespondToPing(c *websocket.Conn, r *http.Request) {
	for {
		_, data, err := c.Read(r.Context())
		if err != nil {
			var websocketError *websocket.CloseError
			if errors.As(err, &websocketError) {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {

				}
				log.Println("Close reason:", websocketError.Reason)
			}
			log.Printf("failed to read message %v: %v", r.RemoteAddr, err)
			return
		}

		log.Println("/dont-respond-to-ping: Got message:", string(data))
	}
}
