package main

import (
	"net/http"

	"github.com/gorilla/mux"

	gohandlers "github.com/gorilla/handlers"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/leaderboard/{bracket:(?:2v2|3v3|rbg)}", app.getLeaderboardByBracket)
	r.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[A-Za-z]+}", app.getCharacter)
	r.HandleFunc("/specs", app.getSpecs)

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/equipment/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacterEquipmentFromBlizzard)
	s.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacterFromBlizzard)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	return ch(r)
}
