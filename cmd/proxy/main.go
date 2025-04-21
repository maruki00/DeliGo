package main

import (
	"context"
	"deligo/cmd/proxy/configs"
	grpc_user "deligo/internal/iam/infra/grpc/user"
	"fmt"
	"log/slog"
	"net/http"

	gruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GateWay(ctx context.Context, cfg *configs.Config, opts []gruntime.ServeMuxOption) (http.Handler, error) {
	mux := gruntime.NewServeMux(opts...)

	userEndPoint := fmt.Sprintf("%s:%s", cfg.GRPCServer.Host, cfg.GRPCServer.Port)

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := grpc_user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userEndPoint, dialOpts); err != nil {
		return nil, err
	}

	return mux, nil
}

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Run request", "http_method", r.Method, "http_url", r.URL)
		h.ServeHTTP(w, r)
	})
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
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		Handler: withLogger(mux),
	}
	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server")

		if err := s.Shutdown(ctx); err != nil {
			slog.Error("could not shutdown the server")
		}
	}()

	fmt.Printf("Starting Server on %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start: ", err)
	}

}
