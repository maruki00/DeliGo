package main

import (
	"context"
	"delivery/cmd/user/configs"
	"delivery/internal/user/app"
	grpc_user "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
)

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Run request", "http_method", r.Method, "http_url", r.URL)
		h.ServeHTTP(w, r)
	})
}

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

	app, clean, err := app.InitApp(cfg)
	if err != nil {
		clean()
		panic(err)

	}
	res, err := app.UserRepo.Create(ctx, &models.User{
		ID:       "123sdfg",
		Email:    "1234qsdfwd",
		Password: "12344qwe",
		Role:     "fsgsdfg",
	})

	return
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Done: %v\n", res)

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	// gRPC Server.
	address := fmt.Sprintf("%s:%s", cfg.GRPCServer.Host, cfg.GRPCServer.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
	}

	slog.Info("🌏 start server...", "address", address)

	grpc_user.RegisterUserServiceServer(server, app.UserSVC)

	defer func() {
		if err := l.Close(); err != nil {
			slog.Error("failed to close", err, "network", network, "address", address)
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", done)
	}
}
