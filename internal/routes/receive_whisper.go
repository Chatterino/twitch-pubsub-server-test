package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func ReceiveWhisper(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()
	const whisper = `{"type":"MESSAGE","data":{"topic":"whispers.11148817","message":"{\"type\":\"whisper_received\",\"data\":\"{\\\"message_id\\\":\\\"e0e48a09-73f4-41e1-b92a-685776084467\\\",\\\"id\\\":9042,\\\"thread_id\\\":\\\"11148817_82008718\\\",\\\"body\\\":\\\"me Kappa\\\",\\\"sent_ts\\\":1650113146,\\\"from_id\\\":82008718,\\\"tags\\\":{\\\"login\\\":\\\"pajbot\\\",\\\"display_name\\\":\\\"pajbot\\\",\\\"color\\\":\\\"#2E8B57\\\",\\\"emotes\\\":[{\\\"emote_id\\\":\\\"25\\\",\\\"start\\\":3,\\\"end\\\":7}],\\\"badges\\\":[]},\\\"recipient\\\":{\\\"id\\\":11148817,\\\"username\\\":\\\"pajlada\\\",\\\"display_name\\\":\\\"pajlada\\\",\\\"color\\\":\\\"#CC44FF\\\"}}\",\"data_object\":{\"message_id\":\"e0e48a09-73f4-41e1-b92a-685776084467\",\"id\":9042,\"thread_id\":\"11148817_82008718\",\"body\":\"me Kappa\",\"sent_ts\":1650113146,\"from_id\":82008718,\"tags\":{\"login\":\"pajbot\",\"display_name\":\"pajbot\",\"color\":\"#2E8B57\",\"emotes\":[{\"emote_id\":\"25\",\"start\":3,\"end\":7}],\"badges\":[]},\"recipient\":{\"id\":11148817,\"username\":\"pajlada\",\"display_name\":\"pajlada\",\"color\":\"#CC44FF\"}}}"}}`

	time.AfterFunc(25*time.Millisecond, func() {
		if err := c.Write(r.Context(), websocket.MessageText, []byte(whisper)); err != nil {
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
