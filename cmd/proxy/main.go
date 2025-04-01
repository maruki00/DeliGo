package main

import (
	"context"
	"delivery/cmd/proxy/configs"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"fmt"
	"log/slog"
	"net/http"

	gruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GateWay(ctx context.Context, cfg *configs.Config, opts []gruntime.ServeMuxOption) (http.Handler, error) {
	mux := gruntime.NewServeMux(opts...)

	userEndPoint := fmt.Sprintf("%s:%s", cfg.UserGRPC.Host, cfg.UserGRPC.Port)

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := user_grpc.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userEndPoint, dialOpts); err != nil {
		return nil, err
	}

	return mux, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	cfg, err := configs.GetConfig()
	if err != nil {
		panic("error getting config: " + err.Error())
	}

	gw, err := GateWay(ctx, cfg, nil)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gw)

	s := http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
	}
	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server")

		if err := s.Shutdown(ctx); err != nil {
			slog.Error("could not shutdown the server")
		}
	}()
	if err := s.ListenAndServe(); err != nil {
		slog.Error("error ", err)
	}

}
