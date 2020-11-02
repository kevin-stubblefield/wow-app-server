package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type config struct {
	blizzardClientId     string
	blizzardClientSecret string
}

type application struct {
	config
	infoLog   *log.Logger
	errorLog  *log.Logger
	client    *http.Client
	cache     cache.Cache
	wowApiUrl string
}

func main() {
	addr := flag.String("addr", ":3000", "Http network address")
	blizzardClientId := flag.String("blizClientId", "", "Client ID for Blizzard API access")
	blizzardClientSecret := flag.String("blizClientSecret", "", "Client secret for Blizzard API access")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := &config{
		blizzardClientId:     *blizzardClientId,
		blizzardClientSecret: *blizzardClientSecret,
	}

	client := &http.Client{Timeout: 10 * time.Second}

	c := cache.New(24*7*time.Hour, 24*8*time.Hour)

	app := &application{
		config:    *cfg,
		infoLog:   infoLog,
		errorLog:  errorLog,
		client:    client,
		cache:     *c,
		wowApiUrl: "https://us.api.blizzard.com/",
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}
