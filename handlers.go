package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"stubblefield.io/wow-leaderboard-api/structs"
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

	resp, err := app.client.Do(req)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.serverError(w, err)
	}

	leaderboard := structs.Leaderboard{}
	err = json.Unmarshal(body, &leaderboard)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte(leaderboard.Entries[0].Character.Name))
}
