# Changelog

## v1.0.0

 - New route `/dont-respond-to-ping`  
   Doesn't respond to pings.
 - New route `/disconnect-client-after-1s`  
   Disconnects the client after 1 second.
 - New route `/receive-whisper`  
   25ms after client connects, the server sends a message as if the client received a whisper.
 - New route `/moderator-actions-user-banned`  
   25ms after client connects, the server sends a message as if the client saw someone in a channel get banned, and they had moderator access to see it.
 - New route `/authentication-required`  
   Route that throws the client out if they listen to a topic without an auth token that's exactly `xD`.
