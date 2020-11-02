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
	infoLog  *log.Logger
	errorLog *log.Logger
	client   *BlizzardClient
}

func main() {
	addr := flag.String("addr", ":3000", "Http network address")
	blizzardClientId := flag.String("blizClientId", "", "Client ID for Blizzard API access")
	blizzardClientSecret := flag.String("blizClientSecret", "", "Client secret for Blizzard API access")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	c := cache.New(24*7*time.Hour, 24*8*time.Hour)

	client := &BlizzardClient{
		cache:                *c,
		wowApiUrl:            "https://us.api.blizzard.com/",
		blizzardClientId:     *blizzardClientId,
		blizzardClientSecret: *blizzardClientSecret,
	}

	client.Timeout = 10 * time.Second

	app := &application{
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
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}
