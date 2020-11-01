package main

import (
	"fmt"
	"net/http"
)

const apiUrl = "https://us.api.blizzard.com/"

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	token := app.getToken()

	pvpSeason := "27"
	pvpLeaderboard := "3v3"

	endpoint := fmt.Sprintf("data/wow/pvp-season/%s/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason, pvpLeaderboard, token.AccessToken)

	req, err := http.NewRequest(http.MethodGet, apiUrl+endpoint, nil)
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
