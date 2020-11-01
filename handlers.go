package main

import (
	"fmt"
	"net/http"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	token := app.getToken()

	pvpSeason := "27"
	pvpLeaderboard := "3v3"

	reqUrl := fmt.Sprintf("https://us.api.blizzard.com/data/wow/pvp-season/%s/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason, pvpLeaderboard, token.AccessToken)

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

	body, err := app.getJSONResponse(req)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
