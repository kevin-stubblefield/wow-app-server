package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 25 {
		limit = 10
	}

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	pvpSeason, err := app.client.getCurrentPvPSeason(token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	leaderboard, err := app.client.getLeaderboardData(pvpSeason.CurrentSeason.Id, pvpBracket, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	results := leaderboard.Entries[offset : offset+limit]

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (app *application) getCharacterEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	equipment, err := app.client.getEquipmentData(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipment)
}

func (app *application) getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	summary, err := app.client.getCharacterSummary(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
