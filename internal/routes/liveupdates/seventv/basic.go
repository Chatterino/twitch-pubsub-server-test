package seventv

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

const (
	emoteSetOldId = "60b39e943e203cc169dfc106"
	emoteSetNewId = "60bca831e7ecd2f892c9b9ab"

	modifyEmoteSetFormatString string = `{"op":0,"t":1667059799105,"d":{"type":"emote_set.update","body":{"id":"%s","kind":3,"actor":{"id":"60b39e943e203cc169dfc106","username":"nerixyz","display_name":"nerixyz","style":{"color":1567670,"paint":null},"biography":"","editors":[{"id":"60b39ec4a4f2ced7f03d92b2","permissions":17,"visible":true,"added_at":1657657522941},{"id":"60be041bd8e231199fa8439d","permissions":17,"visible":true,"added_at":1657657522941},{"id":"60f6159bf7fdd1860acf305a","permissions":17,"visible":true,"added_at":1657657522941}],"roles":["60b3f1ea886e63449c5263b1","62b48deb791a15a25c2a0354"]},"pushed":[{"key":"emotes","index":224,"type":"object","value":{"actor_id":"60b39e943e203cc169dfc106","data":{"animated":false,"flags":0,"host":{"files":[{"format":"AVIF","frame_count":1,"height":32,"name":"1x.avif","size":1504,"static_name":"1x","width":32},{"format":"WEBP","frame_count":1,"height":32,"name":"1x.webp","size":1052,"static_name":"1x","width":32},{"format":"AVIF","frame_count":1,"height":64,"name":"2x.avif","size":2450,"static_name":"2x","width":64},{"format":"WEBP","frame_count":1,"height":64,"name":"2x.webp","size":2462,"static_name":"2x","width":64},{"format":"AVIF","frame_count":1,"height":96,"name":"3x.avif","size":3644,"static_name":"3x","width":96},{"format":"WEBP","frame_count":1,"height":96,"name":"3x.webp","size":4334,"static_name":"3x","width":96},{"format":"AVIF","frame_count":1,"height":128,"name":"4x.avif","size":4634,"static_name":"4x","width":128},{"format":"WEBP","frame_count":1,"height":128,"name":"4x.webp","size":6122,"static_name":"4x","width":128}],"url":"//cdn.7tv.app/emote/621d13967cc2d4e1953838ed"},"id":"621d13967cc2d4e1953838ed","lifecycle":3,"listed":true,"name":"Chatterinoge","owner":{"avatar_url":"//cdn.7tv.app/pp/60ae3e98b2ecb0150535c6b7/4aa1786cec024098be20d7b0683bae72","connections":null,"display_name":"gempir","id":"60ae3e98b2ecb0150535c6b7","roles":["6076a86b09a4c63a38ebe801"],"style":{"color":16755200,"paint":null},"username":"gempir"},"tags":["okayge","chatterino"]},"flags":0,"id":"621d13967cc2d4e1953838ed","name":"Chatterinoge","timestamp":1667059799105}}],"pulled":[{"key":"emotes","index":224,"type":"object","old_value":{"actor_id":"60b39e943e203cc169dfc106","flags":0,"id":"621d13967cc2d4e1953838ed","name":"Chatterino","timestamp":-62135596800000},"value":null}],"updated":[{"key":"emotes","index":224,"type":"object","old_value":{"actor_id":"60b39e943e203cc169dfc106","flags":0,"id":"621d13967cc2d4e1953838ed","name":"Chatterinoge","timestamp":1667059799105},"value":{"actor_id":"60b39e943e203cc169dfc106","data":{"animated":false,"flags":0,"host":{"files":[{"format":"AVIF","frame_count":1,"height":32,"name":"1x.avif","size":1504,"static_name":"1x","width":32},{"format":"WEBP","frame_count":1,"height":32,"name":"1x.webp","size":1052,"static_name":"1x","width":32},{"format":"AVIF","frame_count":1,"height":64,"name":"2x.avif","size":2450,"static_name":"2x","width":64},{"format":"WEBP","frame_count":1,"height":64,"name":"2x.webp","size":2462,"static_name":"2x","width":64},{"format":"AVIF","frame_count":1,"height":96,"name":"3x.avif","size":3644,"static_name":"3x","width":96},{"format":"WEBP","frame_count":1,"height":96,"name":"3x.webp","size":4334,"static_name":"3x","width":96},{"format":"AVIF","frame_count":1,"height":128,"name":"4x.avif","size":4634,"static_name":"4x","width":128},{"format":"WEBP","frame_count":1,"height":128,"name":"4x.webp","size":6122,"static_name":"4x","width":128}],"url":"//cdn.7tv.app/emote/621d13967cc2d4e1953838ed"},"id":"621d13967cc2d4e1953838ed","lifecycle":3,"listed":true,"name":"Chatterinoge","owner":{"avatar_url":"//cdn.7tv.app/pp/60ae3e98b2ecb0150535c6b7/4aa1786cec024098be20d7b0683bae72","connections":null,"display_name":"gempir","id":"60ae3e98b2ecb0150535c6b7","roles":["6076a86b09a4c63a38ebe801"],"style":{"color":16755200,"paint":null},"username":"gempir"},"tags":["okayge","chatterino"]},"flags":0,"id":"621d13967cc2d4e1953838ed","name":"Chatterino","timestamp":1667059799105}}]}}}`

	updateUserConnection string = `{"op":0,"t":1667060748118,"d":{"type":"user.update","body":{"id":"%s","kind":1,"actor":{"id":"60b39e943e203cc169dfc106","username":"nerixyz","display_name":"nerixyz","style":{"color":1567670,"paint":null},"biography":"","editors":[{"id":"60b39ec4a4f2ced7f03d92b2","permissions":17,"visible":true,"added_at":1657657522941},{"id":"60be041bd8e231199fa8439d","permissions":17,"visible":true,"added_at":1657657522941},{"id":"60f6159bf7fdd1860acf305a","permissions":17,"visible":true,"added_at":1657657522941}],"roles":["60b3f1ea886e63449c5263b1","62b48deb791a15a25c2a0354"],"connections":null},"updated":[{"key":"connections","index":0,"nested":true,"type":"","value":[{"index":null,"key":"emote_set","old_value":{"capacity":300,"id":"%s","immutable":false,"name":"test","owner":{"avatar_url":"//static-cdn.jtvnw.net/jtv_user_pictures/e065218b-49df-459d-afd3-c6557870f551-profile_image-70x70.png","connections":null,"display_name":"nerixyz","id":"60b39e943e203cc169dfc106","roles":["60b3f1ea886e63449c5263b1","62b48deb791a15a25c2a0354"],"style":{"color":1567670,"paint":null},"username":"nerixyz"},"privileged":false,"tags":[]},"type":"object","value":{"capacity":300,"id":"%s","immutable":false,"name":"nerixyz's Emotes","owner":{"avatar_url":"//static-cdn.jtvnw.net/jtv_user_pictures/e065218b-49df-459d-afd3-c6557870f551-profile_image-70x70.png","connections":null,"display_name":"nerixyz","id":"60b39e943e203cc169dfc106","roles":["60b3f1ea886e63449c5263b1","62b48deb791a15a25c2a0354"],"style":{"color":1567670,"paint":null},"username":"nerixyz"},"privileged":false,"tags":[]}}]}]}}}`

	helloHb1s string = `{"op":1,"t":1667063289122,"d":{"heartbeat_interval":1000,"session_id":"foo"}}`
)

func AllEvents(c *websocket.Conn, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	if err := c.Write(ctx, websocket.MessageText, []byte(helloHb1s)); err != nil {
		log.Println("Failed to send", err)
		return
	}

	for {
		msg, ok := DefaultRead(ctx, c, r)
		if !ok {
			break
		}

		switch msg.Op {
		case OpcodeSubscribe:
			payload, err := ConvertMessage[SubscribePayload](msg)
			if err != nil {
				log.Println("Bad message", err)
				break
			}
			switch payload.Data.Type {
			case "emote_set.update":
				formatted := []byte(fmt.Sprintf(modifyEmoteSetFormatString, payload.Data.Condition["object_id"]))
				if err := c.Write(ctx, websocket.MessageText, formatted); err != nil {
					log.Println("Failed to send", err)
					break
				}
			case "user.update":
				formatted := []byte(fmt.Sprintf(updateUserConnection, payload.Data.Condition["object_id"], emoteSetOldId, emoteSetNewId))
				if err := c.Write(ctx, websocket.MessageText, formatted); err != nil {
					log.Println("Failed to send", err)
					break
				}
			default:
				log.Println("Unsupported type")
			}

		default:
			log.Println("Unhandled message:", msg)
		}
	}
}
