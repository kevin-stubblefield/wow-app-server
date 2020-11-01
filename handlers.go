package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"stubblefield.io/wow-leaderboard-api/structs"
)

func (app *application) getToken(w http.ResponseWriter, r *http.Request) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, "Unable to create token request", http.StatusInternalServerError)
		app.errorLog.Println("Unable to create token request", err)
	}

	req.SetBasicAuth(app.blizzardClientId, app.blizzardClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.client.Do(req)
	if err != nil {
		http.Error(w, "Token request failed", http.StatusInternalServerError)
		app.errorLog.Println("Token request failed", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed reading token body", http.StatusInternalServerError)
		app.errorLog.Println("Failed reading token body", err)
	}

	token := structs.AuthToken{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		http.Error(w, "Failed parsing token body", http.StatusInternalServerError)
		app.errorLog.Println("Failed parsing token body", err)
	}

	app.infoLog.Println(token)
	w.Write([]byte(token.AccessToken))
}

func (app *application) authCallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
