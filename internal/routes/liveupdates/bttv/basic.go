package bttv

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

const (
	addEmoteFormatString    = `{"data":{"channel":"%s","emote":{"channel":"emjaykae","code":"PepePls","id":"55898e122612142e6aaa935b","imageType":"gif","url":"https://cdn.betterttv.net/emote/55898e122612142e6aaa935b/1x","urlTemplate":"https://cdn.betterttv.net/emote/55898e122612142e6aaa935b/{{image}}","user":{"displayName":"EmJayKae","id":"5537fb2b236a1aa17a9970df","name":"emjaykae","providerId":"23473656"}}},"name":"emote_create"}`
	removeEmoteFormstString = `{"data":{"channel":"%s","emoteId":"55898e122612142e6aaa935b"},"name":"emote_delete"}`
)

func AllEvents(c *websocket.Conn, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	for {
		msg, ok := DefaultRead(ctx, c, r)
		if !ok {
			break
		}

		switch msg.Name {
		case "join_channel":
			payload, err := ConvertMessage[JoinPayload](msg)
			if err != nil {
				log.Println("Bad message", err)
				break
			}
			formatted := []byte(fmt.Sprintf(addEmoteFormatString, payload.Data.Name))
			if err := c.Write(ctx, websocket.MessageText, formatted); err != nil {
				log.Println("Failed to send", err)
				break
			}
			formatted = []byte(fmt.Sprintf(removeEmoteFormstString, payload.Data.Name))
			if err := c.Write(ctx, websocket.MessageText, formatted); err != nil {
				log.Println("Failed to send", err)
				// break
			}

		default:
			log.Println("Unhandled message:", msg)
		}
	}
}
