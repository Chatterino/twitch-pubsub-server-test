package main

import (
	"log"
	"net/http"

	"github.com/Chatterino/twitch-pubsub-server-test/internal/routes"
	"github.com/Chatterino/twitch-pubsub-server-test/internal/routes/liveupdates"
	"github.com/Chatterino/twitch-pubsub-server-test/internal/routes/liveupdates/seventv"
	"nhooyr.io/websocket"
)

type server struct {
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("Error upgrading websocket connection:", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	switch r.URL.Path {
	case "/dont-respond-to-ping":
		routes.DontRespondToPing(c, r)
	case "/disconnect-client-after-1s":
		routes.RandomDisconnect(c, r)
	case "/receive-whisper":
		routes.ReceiveWhisper(c, r)
	case "/moderator-actions-user-banned":
		routes.ModeratorActionsUserBanned(c, r)
	case "/authentication-required":
		routes.AuthenticationRequired(c, r)
	case "/automod-held":
		routes.AutomodHeld(c, r)
	case "/liveupdates/sub-unsub":
		liveupdates.BasicSubUnsub(c, r)
	case "/liveupdates/seventv/all-events":
		seventv.AllEvents(c, r)
	case "/liveupdates/seventv/no-heartbeat":
		seventv.NoHeartbeat(c, r)
	default:
		routes.Default(c, r)
	}
}
