package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"docker-grpc-service-tutorial/server"

	"github.com/pkg/errors"
)

func run(log *log.Logger) error {
	port := 4040
	log.Println("main: Initializing GRPC server")
	defer log.Println("main: Completed")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	serverErrors := make(chan error, 1)

	server, err := server.NewServer(port)
	if err != nil {
		return errors.Wrap(err, "running server")
	}

	go func() {
		log.Printf("main: GRPC server listening on port %d", port)
		serverErrors <- server.Serve()
	}()

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main: %v: Start shutdown", sig)
		server.GracefulStop()
	}

	return nil
}

func main() {
	log := log.New(os.Stdout, "GRPC SERVER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}
