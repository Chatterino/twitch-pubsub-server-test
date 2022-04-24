FROM golang:1.18 AS build
ADD . /src
RUN cd /src/cmd/server && GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags '-extldflags "-static"'

FROM alpine:latest
WORKDIR /app
COPY --from=build /src/cmd/server/server /app/
RUN apk add --no-cache ca-certificates
CMD ["./server"]
