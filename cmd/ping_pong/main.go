package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	ping_pong_v1 "github.com/meo-s/grpc-ping-pong/proto/api/ping_pong/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	ping_pong_v1.RegisterPingPongRpcServer(grpcServer, &ping_pong_v1.PingPongRpcServerImpl{})

	lis, err := net.Listen("tcp", "[::]:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	wg := sync.WaitGroup{}
	cx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		select {
		case <-interrupt:
			grpcServer.Stop()
		case <-cx.Done():
			break
		}
	}()

	defer wg.Wait()
	defer cancel()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
