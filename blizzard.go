package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/patrickmn/go-cache"
	"stubblefield.io/wow-leaderboard-api/data"
)

type BlizzardClient struct {
	http.Client
	cache                cache.Cache
	wowApiUrl            string
	blizzardClientId     string
	blizzardClientSecret string
}

func (client *BlizzardClient) getJSONResponse(req *http.Request, endpoint string, cacheRequest bool) ([]byte, error) {
	cacheValue, found := client.cache.Get(endpoint)
	if found {
		return cacheValue.([]byte), nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if cacheRequest {
		client.cache.Set(endpoint, body, cache.DefaultExpiration)
	}

	return body, nil
}

func (client *BlizzardClient) getToken() (*data.AuthToken, error) {
	reqBody := url.Values{}
	reqBody.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", strings.NewReader(reqBody.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.blizzardClientId, client.blizzardClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := client.getJSONResponse(req, "oauth/token", false)
	if err != nil {
		return nil, err
	}

	token := &data.AuthToken{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (client *BlizzardClient) getCurrentPvPSeason(token string) (*data.SeasonIndex, error) {
	endpoint := fmt.Sprintf("data/wow/pvp-season/index?namespace=dynamic-us&locale=en_US&access_token=%s", token)

	req, err := http.NewRequest(http.MethodGet, client.wowApiUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := client.getJSONResponse(req, endpoint, true)
	if err != nil {
		return nil, err
	}

	pvpSeason := &data.SeasonIndex{}

	err = json.Unmarshal(body, &pvpSeason)
	if err != nil {
		return nil, err
	}

	return pvpSeason, nil
}
