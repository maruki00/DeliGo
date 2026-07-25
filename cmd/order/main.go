package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maruki00/deligo/internal/order"
	"github.com/maruki00/deligo/internal/order/config"
	pkgConfigs "github.com/maruki00/deligo/pkg/config"
)

func main() {

	var cfgPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "config path ")
	flag.Parse()

	cfg, err := pkgConfigs.LoadCfg[config.Config](cfgPath)
	if err != nil {
		log.Fatal("error parse the config")
	}

	bgCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, clean, err := order.Init(cfg)
	if err != nil {
		clean()
		log.Fatal(err)
	}
	defer clean()
	addr := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: app.HTTPServer,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server Listen error variant occurred: %s\n", err)
		}
	}()
	log.Printf("Server online listening securely on addr %s", addr)

	go func() {
		app.Consumer.Start(bgCtx)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal captured. Halting microservice operational tasks safely...")

	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to crash unexpectedly during graceful shutdown loop: %v", err)
	}

	log.Println("Order Service cleanly shut down.")
}
