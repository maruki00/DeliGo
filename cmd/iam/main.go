package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/maruki00/deligo/cmd/iam/configs"
	"github.com/maruki00/deligo/internal/iam/app"
	grpc_user "github.com/maruki00/deligo/internal/iam/infra/grpc/user"

	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if v, ok := req.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return handler(ctx, req)
}

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

	server := grpc.NewServer(
		grpc.UnaryInterceptor(validationInterceptor),
	)

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

	grpc_user.RegisterUserServiceServer(server, app.UserServerSvc)

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
