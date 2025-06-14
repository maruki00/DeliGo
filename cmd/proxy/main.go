package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/maruki00/github.com/maruki00/deligo/cmd/proxy/configs"
	grpc_user "github.com/maruki00/github.com/maruki00/deligo/internal/iam/infra/grpc/user"

	gruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func GateWay(ctx context.Context, cfg *configs.Config, opts []gruntime.ServeMuxOption) (http.Handler, error) {
	start := time.Now()
	defer func() {
		slog.Info("Gateway setup completed", "duration", time.Since(start))
	}()

	mux := gruntime.NewServeMux(opts...)
	fmt.Printf("%s:%s", cfg.IAMGRPC.Host, cfg.IAMGRPC.Port)
	userEndPoint := "0.0.0.0:9001" // fmt.Sprintf("%s:%s", cfg.IAMGRPC.Host, cfg.IAMGRPC.Port)
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientInterceptor),
	}
	if err := grpc_user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userEndPoint, dialOpts); err != nil {
		return nil, err
	}
	return mux, nil
}

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	slog.Info("gRPC request completed",
		"method", method,
		"duration", time.Since(start).String(),
		"error", errMsg)
	return err
}

// Wrap HTTP handler with timing and logging
func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := newResponseWriter(w)
		h.ServeHTTP(rw, r)
		slog.Info("HTTP request completed",
			"method", r.Method,
			"url", r.URL.String(),
			"status", rw.statusCode,
			"duration", time.Since(start).String())
	})
}

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CustomMetadata(ctx context.Context, req *http.Request) metadata.MD {
	return metadata.Pairs("x-raw-query", req.URL.RawQuery)
}

func main() {
	totalStart := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration
	configStart := time.Now()
	cfg, err := configs.GetConfig()
	if err != nil {
		panic("error getting config: " + err.Error())
	}
	slog.Info("Config loaded", "duration", time.Since(configStart))

	gwStart := time.Now()
	gw, err := GateWay(ctx, cfg, []gruntime.ServeMuxOption{gruntime.WithMetadata(CustomMetadata)})
	if err != nil {
		panic(err)
	}
	slog.Info("Gateway setup completed", "duration", time.Since(gwStart))

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	s := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		Handler: withLogger(mux),
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		shutdownStart := time.Now()
		slog.Info("Shutting down server")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			slog.Error("could not shutdown the server", "error", err)
		}
		slog.Info("Server shutdown completed", "duration", time.Since(shutdownStart))
	}()

	slog.Info("Server initialization completed",
		"total_setup_time", time.Since(totalStart),
		"listen_address", s.Addr)

	fmt.Printf("Starting Server on %s:%s\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
	}
}
