package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"stubblefield.io/wow-leaderboard-api/data"
)

type BlizzardClient struct {
	http.Client
	wowApiUrl            string
	blizzardClientId     string
	blizzardClientSecret string
}

func (client *BlizzardClient) getJSONResponse(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
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

	body, err := client.getJSONResponse(req)
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

	body, err := client.getJSONResponse(req)
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

func (client *BlizzardClient) getLeaderboardData(pvpSeason int, pvpBracket, token string) (*data.Leaderboard, error) {
	endpoint := fmt.Sprintf("data/wow/pvp-season/%d/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason, pvpBracket, token)

	req, err := http.NewRequest(http.MethodGet, client.wowApiUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	leaderboard := &data.Leaderboard{}

	json.NewDecoder(resp.Body).Decode(&leaderboard)

	return leaderboard, nil
}

func (client *BlizzardClient) getEquipmentData(realmSlug, character, token string) (*data.Equipment, error) {
	endpoint := fmt.Sprintf("profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token)

	req, err := http.NewRequest(http.MethodGet, client.wowApiUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := client.getJSONResponse(req)
	if err != nil {
		return nil, err
	}

	equipment := &data.Equipment{}

	err = json.Unmarshal(body, &equipment)
	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (client *BlizzardClient) getCharacterSummary(realmSlug, character, token string) (*data.CharacterSummary, error) {
	realmSlug = strings.ToLower(realmSlug)
	character = strings.ToLower(character)
	endpoint := fmt.Sprintf("profile/wow/character/%s/%s?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token)

	req, err := http.NewRequest(http.MethodGet, client.wowApiUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	summary := &data.CharacterSummary{}

	json.NewDecoder(resp.Body).Decode(&summary)

	return summary, nil
}
