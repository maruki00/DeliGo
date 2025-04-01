package main

import (
	"context"
	"delivery/cmd/user/configs"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
)

func main() {

	_, err := maxprocs.Set()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := configs.GetConfig()
	if err != nil {
		cancel()
		panic(err)
	}

	// server := grpc.NewServer()
	// go func() {
	// 	server.GracefulStop()
	// 	<-ctx.Done()
	// }()

	// //Grpc Server
	// protocol := "tcp"
	// addr := fmt.Sprintf("%s:%s", cfg.GRPCServer.Host, cfg.GRPCServer.Port)

	// listner, err := net.Listen(protocol, addr)
	// if err != nil {
	// 	cancel()
	// 	panic(err)
	// }
	// fmt.Printf("Server Starting on %s:%s\n", cfg.GRPCServer.Host, cfg.GRPCServer.Port)
	// server.Serve(listner)

	// Force port 9001
	cfg.GRPCServer.Host = "0.0.0.0"
	cfg.GRPCServer.Port = "9001"

	// Create gRPC server
	server := grpc.NewServer()

	// Start listener
	addr := fmt.Sprintf("%s:%s", cfg.GRPCServer.Host, cfg.GRPCServer.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	log.Printf("🚀 gRPC Server is running on %s", addr)

	// Handle graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("Shutting down gRPC server...")
		server.GracefulStop()
	}()

	// Start gRPC server
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		slog.Info("signal.Notify ", v)
	case done := <-ctx.Done():
		slog.Info("done ", done)
	}
}
