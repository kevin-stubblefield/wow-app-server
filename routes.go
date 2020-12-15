package main

import (
	"net/http"

	"github.com/gorilla/mux"

	gohandlers "github.com/gorilla/handlers"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", app.showDashboard)
	r.HandleFunc("/leaderboard/{bracket:(?:2v2|3v3|rbg)}", app.showLeaderboard)
	r.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[A-Za-z]+}", app.showCharacter)
	r.HandleFunc("/specs", app.getSpecs)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacter)

	b := s.PathPrefix("/blizz").Subrouter()
	b.HandleFunc("/equipment/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacterEquipmentFromBlizzard)
	b.HandleFunc("/character/{realmSlug:[a-z-]+}/{character:[a-z]+}", app.getCharacterFromBlizzard)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	return ch(r)
}
