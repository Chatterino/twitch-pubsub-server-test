package routes

import (
	"log"
	"net/http"
	"time"

	"git.kotmisia.pl/Mm2PL/examples"
	"nhooyr.io/websocket"
)

func ModeratorActionsUserBanned(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()
	exampleMessage := examples.PubsubChatModeratorActionsUserBanned()

	time.AfterFunc(25*time.Millisecond, func() {
		if err := c.Write(r.Context(), websocket.MessageText, []byte(exampleMessage)); err != nil {
			log.Println("Error writing response", err)
			return
		}
	})

	for {
		doBreak := defaultHandler(ctx, c, r)
		if doBreak {
			break
		}
	}
}
