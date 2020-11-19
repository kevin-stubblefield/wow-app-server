package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"stubblefield.io/wow-leaderboard-api/data"
)

// BlizzardClient Custom HTTP client to interact with Blizzard APIs
type BlizzardClient struct {
	http.Client
	BaseURL              string
	BlizzardClientID     string
	BlizzardClientSecret string
}

// GetToken Retrieves token from Blizzard API for authentication
func (client *BlizzardClient) GetToken() (*data.AuthToken, error) {
	reqBody := url.Values{}
	reqBody.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", strings.NewReader(reqBody.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.BlizzardClientID, client.BlizzardClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	token := &data.AuthToken{}

	json.NewDecoder(resp.Body).Decode(&token)

	return token, nil
}

// GetCurrentPvPSeason Retrieves the current pvp season from the Blizzard API
func (client *BlizzardClient) GetCurrentPvPSeason(token string) (*data.SeasonIndex, error) {
	endpoint := fmt.Sprintf("data/wow/pvp-season/index?namespace=dynamic-us&locale=en_US&access_token=%s", token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pvpSeason := &data.SeasonIndex{}

	json.NewDecoder(resp.Body).Decode(&pvpSeason)

	return pvpSeason, nil
}

// GetLeaderboardData Retrieves the leaderboard for the given season and bracket from the Blizzard API
func (client *BlizzardClient) GetLeaderboardData(pvpSeason int, pvpBracket, token string) (*data.Leaderboard, error) {
	endpoint := fmt.Sprintf("data/wow/pvp-season/%d/pvp-leaderboard/%s?namespace=dynamic-us&locale=en_US&access_token=%s", pvpSeason, pvpBracket, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	leaderboard := &data.Leaderboard{}

	json.NewDecoder(resp.Body).Decode(&leaderboard)

	return leaderboard, nil
}

// GetCharacterSummary Retrieves character data for the specified character from Blizzard API
func (client *BlizzardClient) GetCharacterSummary(realmSlug, character, token string) (*data.CharacterSummary, error) {
	realmSlug = strings.ToLower(realmSlug)
	character = strings.ToLower(character)
	endpoint := fmt.Sprintf("profile/wow/character/%s/%s?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	summary := &data.CharacterSummary{}

	json.NewDecoder(resp.Body).Decode(&summary)

	return summary, nil
}

// GetCharacterSpecs Retrieves the playable specializations from the Blizzard API
func (client *BlizzardClient) GetCharacterSpecs(token string) (*data.SpecializationIndex, error) {
	endpoint := fmt.Sprintf("data/wow/playable-specialization/index?namespace=static-us&locale=en_US&access_token=%s", token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	specs := &data.SpecializationIndex{}
	json.NewDecoder(resp.Body).Decode(&specs)
	return specs, nil
}

// GetSpecIcon Retrieves the icon for the given spec ID from the Blizzard API
func (client *BlizzardClient) GetSpecIcon(specID int, token string) (*data.Specialization, io.ReadCloser, error) {
	endpoint := fmt.Sprintf("data/wow/playable-specialization/%d?namespace=static-us&locale=en_US&access_token=%s", specID, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	spec := &data.Specialization{}
	json.NewDecoder(resp.Body).Decode(&spec)

	iconResp, err := client.Get(fmt.Sprintf("%s&locale=en_US&access_token=%s", spec.Media.Key.Href, token))
	if err != nil {
		return nil, nil, err
	}

	icon := &data.SpecIcon{}
	json.NewDecoder(iconResp.Body).Decode(&icon)

	imageResp, err := client.Get(icon.Assets[0].Value)
	if err != nil {
		return nil, nil, err
	}

	return spec, imageResp.Body, nil
}

// GetCharacterClasses Retrieves the playable classes from the Blizzard API
func (client *BlizzardClient) GetCharacterClasses(token string) (*data.ClassIndex, error) {
	endpoint := fmt.Sprintf("data/wow/playable-class/index?namespace=static-us&locale=en_US&access_token=%s", token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	classes := &data.ClassIndex{}
	json.NewDecoder(resp.Body).Decode(&classes)
	return classes, nil
}

// GetClassIcon Retrieves the icon for the given class from the Blizzard API
func (client *BlizzardClient) GetClassIcon(classID int, token string) (io.ReadCloser, error) {
	endpoint := fmt.Sprintf("data/wow/media/playable-class/%d?namespace=static-us&locale=en_US&access_token=%s", classID, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	icon := &data.ClassIcon{}
	json.NewDecoder(resp.Body).Decode(&icon)

	imageResp, err := client.Get(icon.Assets[0].Value)
	if err != nil {
		return nil, err
	}

	return imageResp.Body, nil
}

// GetCharacterEquipment Retrieves the equipment for the given character from the Blizzard API
func (client *BlizzardClient) GetCharacterEquipment(realmSlug, character, token string) (*data.Equipment, error) {
	realmSlug = strings.ToLower(realmSlug)
	character = strings.ToLower(character)
	endpoint := fmt.Sprintf("profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	equipment := &data.Equipment{}
	json.NewDecoder(resp.Body).Decode(&equipment)

	return equipment, nil
}

// GetCharacterMedia retrieves the media hrefs for the given character from the Blizzard API
func (client *BlizzardClient) GetCharacterMedia(realmSlug, character, token string) (*data.CharacterMedia, error) {
	realmSlug = strings.ToLower(realmSlug)
	character = strings.ToLower(character)
	endpoint := fmt.Sprintf("profile/wow/character/%s/%s/character-media?namespace=profile-us&locale=en_US&access_token=%s", realmSlug, character, token)

	resp, err := client.Get(client.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	media := &data.CharacterMedia{}
	json.NewDecoder(resp.Body).Decode(&media)

	return media, nil
}
