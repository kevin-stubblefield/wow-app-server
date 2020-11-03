package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/leaderboard/{bracket:(?:2v2|3v3|rbg)}", app.getLeaderboardByBracket)
	s.HandleFunc("/equipment/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacterEquipment)
	s.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacter)

	return r
}
