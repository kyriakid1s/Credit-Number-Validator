package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/form/v4"
)

type application struct {
	logger      *slog.Logger
	formDecoder *form.Decoder
}

func main() {
	port := flag.String("port", ":4000", "HTTP network port")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	formDecoder := form.NewDecoder()

	app := &application{
		logger:      logger,
		formDecoder: formDecoder,
	}
	server := &http.Server{
		Addr:     *port,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting Server", slog.String("port", *port))

	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
