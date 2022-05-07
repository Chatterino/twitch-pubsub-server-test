package routes

import (
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

/*
Message held: (Reason: Misogyny)
{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"misogyny\",\"level\":2},\"message\":{\"content\":{\"text\":\"kurwa\",\"fragments\":[{\"text\":\"kurwa\",\"automod\":{\"topics\":{\"bullying\":7,\"identity\":6,\"vulgar\":6}}}]},\"id\":\"fcc4d00a-2e6d-44d5-b3a1-e08dd9d8bab7\",\"sender\":{\"user_id\":\"99631238\",\"login\":\"zneix\",\"display_name\":\"zneix\",\"chat_color\":\"#F97304\",\"badges\":[{\"id\":\"glhf-pledge\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:56:03.784845576Z\"},\"reason_code\":\"\",\"resolver_id\":\"\",\"resolver_login\":\"\",\"status\":\"PENDING\"}}"}}

Message held: (Reason: Swearing)
{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"swearing\",\"level\":4},\"message\":{\"content\":{\"text\":\"big  cock\",\"fragments\":[{\"text\":\"big  cock\",\"automod\":{\"topics\":{\"dating_and_sexting\":5,\"vulgar\":3}}}]},\"id\":\"8d99df28-6ed1-49f2-b036-a1984cc9c11c\",\"sender\":{\"user_id\":\"11148817\",\"login\":\"pajlada\",\"display_name\":\"pajlada\",\"chat_color\":\"#CC44FF\",\"badges\":[{\"id\":\"partner\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:56:47.800759162Z\"},\"reason_code\":\"\",\"resolver_id\":\"\",\"resolver_login\":\"\",\"status\":\"PENDING\"}}"}}

Message held: (Reason: Sex-based terms)
{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"sexwords\",\"level\":3},\"message\":{\"content\":{\"text\":\"vagina\",\"fragments\":[{\"text\":\"vagina\",\"automod\":{\"topics\":{\"dating_and_sexting\":5}}}]},\"id\":\"db92aa89-e9ad-4c5d-9c1f-033732d10988\",\"sender\":{\"user_id\":\"99631238\",\"login\":\"zneix\",\"display_name\":\"zneix\",\"chat_color\":\"#F97304\",\"badges\":[{\"id\":\"glhf-pledge\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:57:21.147063587Z\"},\"reason_code\":\"\",\"resolver_id\":\"\",\"resolver_login\":\"\",\"status\":\"PENDING\"}}"}}

# Big cock denied
{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"swearing\",\"level\":4},\"message\":{\"content\":{\"text\":\"big  cock\",\"fragments\":[{\"text\":\"big  cock\",\"automod\":{\"topics\":{\"dating_and_sexting\":5,\"vulgar\":3}}}]},\"id\":\"8d99df28-6ed1-49f2-b036-a1984cc9c11c\",\"sender\":{\"user_id\":\"11148817\",\"login\":\"pajlada\",\"display_name\":\"pajlada\",\"chat_color\":\"#CC44FF\",\"badges\":[{\"id\":\"partner\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:56:47.800759162Z\"},\"reason_code\":\"\",\"resolver_id\":\"117166826\",\"resolver_login\":\"testaccount_420\",\"status\":\"DENIED\"}}"}}
{"type":"MESSAGE","data":{"topic":"chat_moderator_actions.117166826.117166826","message":"{\"type\":\"channel_terms_action\",\"data\":{\"type\":\"add_blocked_term\",\"id\":\"b1d1593a-60d8-402c-ad3e-fe2cca520751\",\"text\":\"big  cock\",\"requester_id\":\"117166826\",\"requester_login\":\"testaccount_420\",\"channel_id\":\"117166826\",\"expires_at\":\"2022-05-01T11:57:40.012119609Z\",\"updated_at\":\"2022-05-01T10:57:40.012118923Z\",\"from_automod\":true}}"}}

# Vagina allowed
{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"sexwords\",\"level\":3},\"message\":{\"content\":{\"text\":\"vagina\",\"fragments\":[{\"text\":\"vagina\",\"automod\":{\"topics\":{\"dating_and_sexting\":5}}}]},\"id\":\"db92aa89-e9ad-4c5d-9c1f-033732d10988\",\"sender\":{\"user_id\":\"99631238\",\"login\":\"zneix\",\"display_name\":\"zneix\",\"chat_color\":\"#F97304\",\"badges\":[{\"id\":\"glhf-pledge\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:57:21.147063587Z\"},\"reason_code\":\"\",\"resolver_id\":\"117166826\",\"resolver_login\":\"testaccount_420\",\"status\":\"ALLOWED\"}}"}}
{"type":"MESSAGE","data":{"topic":"chat_moderator_actions.117166826.117166826","message":"{\"type\":\"channel_terms_action\",\"data\":{\"type\":\"add_permitted_term\",\"id\":\"2b6952ca-c7e4-49fe-8ea4-d8262ec146c5\",\"text\":\"vagina\",\"requester_id\":\"117166826\",\"requester_login\":\"testaccount_420\",\"channel_id\":\"117166826\",\"expires_at\":\"2022-05-01T11:58:14.189115313Z\",\"updated_at\":\"2022-05-01T10:58:14.189114442Z\",\"from_automod\":true}}"}}
*/

func AutomodHeld(c *websocket.Conn, r *http.Request) {
	ctx := r.Context()
	const whisper = `{"type":"MESSAGE","data":{"topic":"automod-queue.117166826.117166826","message":"{\"type\":\"automod_caught_message\",\"data\":{\"content_classification\":{\"category\":\"misogyny\",\"level\":2},\"message\":{\"content\":{\"text\":\"kurwa\",\"fragments\":[{\"text\":\"kurwa\",\"automod\":{\"topics\":{\"bullying\":7,\"identity\":6,\"vulgar\":6}}}]},\"id\":\"fcc4d00a-2e6d-44d5-b3a1-e08dd9d8bab7\",\"sender\":{\"user_id\":\"99631238\",\"login\":\"zneix\",\"display_name\":\"zneix\",\"chat_color\":\"#F97304\",\"badges\":[{\"id\":\"glhf-pledge\",\"version\":\"1\"}]},\"sent_at\":\"2022-05-01T10:56:03.784845576Z\"},\"reason_code\":\"\",\"resolver_id\":\"\",\"resolver_login\":\"\",\"status\":\"PENDING\"}}"}}`

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
