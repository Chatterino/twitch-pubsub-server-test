package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func ReceiveWhisper(c *websocket.Conn, r *http.Request) {
	const whisper = `{"type":"MESSAGE","data":{"topic":"whispers.11148817","message":"{\"type\":\"whisper_received\",\"data\":\"{\\\"message_id\\\":\\\"e0e48a09-73f4-41e1-b92a-685776084467\\\",\\\"id\\\":9042,\\\"thread_id\\\":\\\"11148817_82008718\\\",\\\"body\\\":\\\"me Kappa\\\",\\\"sent_ts\\\":1650113146,\\\"from_id\\\":82008718,\\\"tags\\\":{\\\"login\\\":\\\"pajbot\\\",\\\"display_name\\\":\\\"pajbot\\\",\\\"color\\\":\\\"#2E8B57\\\",\\\"emotes\\\":[{\\\"emote_id\\\":\\\"25\\\",\\\"start\\\":3,\\\"end\\\":7}],\\\"badges\\\":[]},\\\"recipient\\\":{\\\"id\\\":11148817,\\\"username\\\":\\\"pajlada\\\",\\\"display_name\\\":\\\"pajlada\\\",\\\"color\\\":\\\"#CC44FF\\\"}}\",\"data_object\":{\"message_id\":\"e0e48a09-73f4-41e1-b92a-685776084467\",\"id\":9042,\"thread_id\":\"11148817_82008718\",\"body\":\"me Kappa\",\"sent_ts\":1650113146,\"from_id\":82008718,\"tags\":{\"login\":\"pajbot\",\"display_name\":\"pajbot\",\"color\":\"#2E8B57\",\"emotes\":[{\"emote_id\":\"25\",\"start\":3,\"end\":7}],\"badges\":[]},\"recipient\":{\"id\":11148817,\"username\":\"pajlada\",\"display_name\":\"pajlada\",\"color\":\"#CC44FF\"}}}"}}`

	fmt.Println(whisper)
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

		var msg Message

		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("Failed to unmarshal data", err)
			return
		}

		switch msg.Type {
		case "PING":
			// respond with pong
			if err := c.Write(r.Context(), websocket.MessageText, pongPayload); err != nil {
				log.Println("Error writing response", err)
				return
			}
			if err := c.Write(r.Context(), websocket.MessageText, []byte(whisper)); err != nil {
				log.Println("Error writing response", err)
				return
			}
		}

		log.Println("/: Got message:", msg)
	}
}
