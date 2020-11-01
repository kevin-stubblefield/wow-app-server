package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	blizzardClientId     string
	blizzardClientSecret string
}

type application struct {
	config
	infoLog  *log.Logger
	errorLog *log.Logger
	client   *http.Client
}

func main() {
	addr := flag.String("addr", ":3000", "Http network address")
	blizzardClientId := flag.String("blizClientId", "9539f43fe1784bef9ff62aee95727bb6", "Client ID for Blizzard API access")
	blizzardClientSecret := flag.String("blizClientSecret", "66HF224N67cAaDfkCqqCH2DtmFm5xzNT", "Client secret for Blizzard API access")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := &config{
		blizzardClientId:     *blizzardClientId,
		blizzardClientSecret: *blizzardClientSecret,
	}

	client := &http.Client{Timeout: 10 * time.Second}

	app := &application{
		config:   *cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		client:   client,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
