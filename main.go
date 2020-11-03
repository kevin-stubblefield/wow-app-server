package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"stubblefield.io/wow-leaderboard-api/models/sqlite"
)

type config struct {
	blizzardClientId     string
	blizzardClientSecret string
}

type application struct {
	infoLog     *log.Logger
	errorLog    *log.Logger
	leaderboard sqlite.PvpLeaderboardStore
	client      *BlizzardClient
}

func main() {
	addr := flag.String("addr", ":3000", "Http network address")
	dsn := flag.String("dsn", "./leaderboard.db", "Data Source Name for the database")
	blizzardClientId := flag.String("blizClientId", "", "Client ID for Blizzard API access")
	blizzardClientSecret := flag.String("blizClientSecret", "", "Client secret for Blizzard API access")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	client := &BlizzardClient{
		wowApiUrl:            "https://us.api.blizzard.com/",
		blizzardClientId:     *blizzardClientId,
		blizzardClientSecret: *blizzardClientSecret,
	}

	client.Timeout = 10 * time.Second

	app := &application{
		infoLog:     infoLog,
		errorLog:    errorLog,
		leaderboard: sqlite.PvpLeaderboardStore{DB: db},
		client:      client,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}
