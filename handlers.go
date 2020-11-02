package main

import (
	"fmt"
	"net/http"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
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

	pvpLeaderboard := "3v3"

	endpoint := fmt.Sprintf("data/wow/pvp-season/%d/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason.CurrentSeason.Id, pvpLeaderboard, token.AccessToken)

	req, err := http.NewRequest(http.MethodGet, app.wowApiUrl+endpoint, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

	body, err := app.getJSONResponse(req, endpoint)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
