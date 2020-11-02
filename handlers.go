package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	token, err := app.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	pvpSeason, err := app.getCurrentPvPSeason(token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	endpoint := fmt.Sprintf("data/wow/pvp-season/%d/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason.CurrentSeason.Id, pvpBracket, token.AccessToken)

	req, err := http.NewRequest(http.MethodGet, app.wowApiUrl+endpoint, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

	body, err := app.getJSONResponse(req, endpoint, true)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func (app *application) getCharacterEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	endpoint := fmt.Sprintf("profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token.AccessToken)

	req, err := http.NewRequest(http.MethodGet, app.wowApiUrl+endpoint, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

	body, err := app.getJSONResponse(req, endpoint, true)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
