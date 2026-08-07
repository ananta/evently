package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
  port int
  env string
}

type application struct {
  config config
  logger *slog.Logger
}

func main() {

  var cfg config

  flag.IntVar(&cfg.port, "port", 4000, "API server port")
  flag.StringVar(&cfg.env, "env", "dev", "Application environment")
  flag.Parse()

  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

  app := application {
    config: cfg,
    logger: logger,
    models: data.NewModels(db),
  }

  srv := &http.Server{
    Addr: fmt.Sprintf(":%d", cfg.port),
    Handler: app.routes(),
    WriteTimeout: time.Second * 30,
    ReadTimeout: time.Second * 10,
    IdleTimeout: time.Minute,
    ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
  }

  logger.Info("starting server", "port", srv.Addr)

  err := srv.ListenAndServe()
  logger.Error(err.Error())
  os.Exit(1)


}
