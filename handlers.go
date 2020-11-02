package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

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

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
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
