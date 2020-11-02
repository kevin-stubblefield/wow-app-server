package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/leaderboard/{bracket:(?:2v2|3v3|rbg)}", app.getLeaderboard)
	r.HandleFunc("/equipment/{realmSlug:[a-z]+}/{character:[a-z]+}", app.getCharacterEquipment)

	return r
}
