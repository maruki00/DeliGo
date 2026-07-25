package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maruki00/deligo/internal/user"
	"github.com/maruki00/deligo/internal/user/config"
	pkgConfigs "github.com/maruki00/deligo/pkg/config"
)

func main() {

	var cfgFile string
	flag.StringVar(&cfgFile, "config", "config.yaml", "config path")
	flag.Parse()

	cfg, err := pkgConfigs.LoadCfg[config.Config](cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	app, clean, err := user.Init(cfg)
	if err != nil {
		clean()
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	slog.Info("sever", "started", addr)
	log.Fatal(app.HTTPServer.Run(addr))

	server := &http.Server{Addr: addr, Handler: app.HTTPServer}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Network socket layer terminated unexpectedly: %v", err)
		}
	}()
	log.Printf("Unified IAM and User Profile microservice launched securely on addr: %s", addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
	log.Println("Operational engine shut down safely without resource leaks.")
}
