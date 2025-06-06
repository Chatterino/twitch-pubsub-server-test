# Changelog

## Unreleased

## v1.0.12

- Add `/URL` command to the `/echo` handler. (#126)

## v1.0.11

- Add default TCP listener param to dockerfile. (#124)

## v1.0.10

- Add `/HEADER <name>` command to the `/echo` handler. (#121)

## v1.0.9

- Add optional TCP listener (second argument). (#120)
- Add `/echo` handler. (#120)

## v1.0.8

- Include `server.crt` and `server.key` in releases. (#113)

## v1.0.7

- Only dev changes, like bumping go version etc

## v1.0.6

- Add BTTV live emote update tests. (#33, #34)

## v1.0.5

- Add 7TV EventApi tests. (#30)

## v1.0.4

- Add test for generic pubsub. (#27)

## v1.0.3

- Update docker build dependencies.
- New route `/automod-held`  
  Sends an "automod held" message shorly after client connected.

## v1.0.2

- Update docker build dependencies.
- Add certificate and private key to docker image.
- Only build for Ubuntu

## v1.0.1

- Fix docker build not publishing on tags.

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
