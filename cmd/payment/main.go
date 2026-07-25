package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maruki00/deligo/internal/payment"
	"github.com/maruki00/deligo/internal/payment/config"
	pkgConfigs "github.com/maruki00/deligo/pkg/config"
)

func main() {

	var cfgPath string
	var mgrPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "config path")
	flag.StringVar(&mgrPath, "migration", "", "migrations path")
	flag.Parse()

	if mgrPath != "" {
		//migrations goes here
	}

	cfg, err := pkgConfigs.LoadCfg[config.Config](cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	app, clean, err := payment.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer clean()

	addr := fmt.Sprintf("%s:%s", cfg.HTTPSerer.Host, cfg.HTTPSerer.Port)
	srv := http.Server{
		Addr:    addr,
		Handler: app.HTTPServer,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		app.AMQPClient.Consumer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
