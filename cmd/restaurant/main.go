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

	"github.com/maruki00/deligo/internal/restaurant"
	"github.com/maruki00/deligo/internal/restaurant/config"
	"github.com/maruki00/deligo/internal/restaurant/repositories"
	pkgConfigs "github.com/maruki00/deligo/pkg/config"
)

func main() {

	var cfgPath string
	var mgrPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "config path")
	flag.StringVar(&mgrPath, "migration", "", "migration sql path")
	flag.Parse()

	cfg, err := pkgConfigs.LoadCfg[config.Config](cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if mgrPath != "" {
		repositories.RunDatabaseMigrations(cfg.DB.Dsn, mgrPath)
		os.RemoveAll(mgrPath)
	}

	app, clean, err := restaurant.Iinit(cfg)
	if err != nil {
		clean()
		log.Fatal(err)
	}
	defer clean()

	addr := fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: app.HTTPServer,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("could not start the server at %s, %v", addr, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
