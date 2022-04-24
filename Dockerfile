FROM golang:1.18 AS build
ADD . /src
RUN cd /src/cmd/server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:latest
WORKDIR /app
COPY --from=build /src/cmd/server/server /app/
COPY --from=build /src/cmd/server/server.crt /app/
COPY --from=build /src/cmd/server/server.key /app/
CMD ["./server", "127.0.0.1:9050"]
