package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"stubblefield.io/wow-leaderboard-api/structs"
)

func (app *application) getToken() structs.AuthToken {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		app.errorLog.Println("Unable to create token request", err)
	}

	req.SetBasicAuth(app.blizzardClientId, app.blizzardClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.client.Do(req)
	if err != nil {
		app.errorLog.Println("Token request failed", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.errorLog.Println("Failed reading token body", err)
	}

	token := structs.AuthToken{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		app.errorLog.Println("Failed parsing token body", err)
	}

	return token
}
