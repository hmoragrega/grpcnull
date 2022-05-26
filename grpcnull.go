package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var omitAddressInUseErr bool
	flag.BoolVar(&omitAddressInUseErr, "omit-address-in-use-error", false,
		"Set to true to silently discard the error if the port is already in use")

	if len(os.Args) > 2 {
		_ = flag.CommandLine.Parse(os.Args[2:])
	}

	address := os.Args[1]
	listener, err := net.Listen("tcp", address)
	if err != nil {
		// Log the error always
		log.Println(err)

		// Default to error exit code
		exit := 1
		if omitAddressInUseErr && strings.Contains(err.Error(), "address already in use") {
			exit = 0
		}

		os.Exit(exit)
		return
	}
	server := grpc.NewServer()

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	log.Println("gRPC null listening at", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("gRPC server at %s failed with: %v\n", address, err)
	}
}
