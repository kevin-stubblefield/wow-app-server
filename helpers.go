package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"

	"github.com/patrickmn/go-cache"
	"stubblefield.io/wow-leaderboard-api/structs"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) getJSONResponse(req *http.Request, endpoint string) ([]byte, error) {
	cacheValue, found := app.cache.Get(endpoint)
	if found {
		return cacheValue.([]byte), nil
	}

	resp, err := app.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	app.cache.Set(endpoint, body, cache.DefaultExpiration)

	return body, nil
}

func (app *application) getToken() (*structs.AuthToken, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(app.blizzardClientId, app.blizzardClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	token := &structs.AuthToken{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (app *application) getCurrentPvPSeason(token string) (*structs.SeasonIndex, error) {
	endpoint := fmt.Sprintf("data/wow/pvp-season/index?namespace=dynamic-us&locale=en_US&access_token=%s", token)

	req, err := http.NewRequest(http.MethodGet, app.wowApiUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := app.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pvpSeason := &structs.SeasonIndex{}

	err = json.Unmarshal(body, &pvpSeason)
	if err != nil {
		return nil, err
	}

	return pvpSeason, nil
}
