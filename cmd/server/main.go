package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e1, e2 := run()
	if e1 != nil {
		log.Fatal(e1)
	}
	if e2 != nil {
		log.Fatal(e2)
	}
}

// run starts a http.Server for the passed in address
// with all requests handled by echoServer.
func run() (error, error) {
	if len(os.Args) < 2 {
		return errors.New("please provide an address to listen on as the first argument"), nil
	}

	tlsListener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		return err, nil
	}

	log.Printf("listening on https://%v", tlsListener.Addr())

	tlsServer := &http.Server{
		Handler:      server{},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	certFile := "server.crt"
	keyFile := "server.key"
	errc := make(chan error, 1)
	go func() {
		errc <- tlsServer.ServeTLS(tlsListener, certFile, keyFile)
	}()

	var tcpServer *http.Server
	if len(os.Args) >= 3 {
		tcpListener, err := net.Listen("tcp", os.Args[2])
		if err != nil {
			return err, nil
		}

		log.Printf("listening on http://%v", tcpListener.Addr())

		tcpServer = &http.Server{
			Handler:      server{},
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		go func() {
			errc <- tcpServer.Serve(tcpListener)
		}()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var tcpErr error
	if tcpServer != nil {
		tcpErr = tcpServer.Shutdown(ctx)
	}
	return tlsServer.Shutdown(ctx), tcpErr
}
