package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"stubblefield.io/wow-leaderboard-api/data"
)

func (app *application) getLeaderboardByBracket(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	leaderboard, err := app.leaderboard.FetchAllByBracket(pvpBracket, "", "")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 25 {
		limit = 10
	}

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	pvpSeason, err := app.client.getCurrentPvPSeason(token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	leaderboard, err := app.client.getLeaderboardData(pvpSeason.CurrentSeason.Id, pvpBracket, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	result := &data.Leaderboard{}

	result.Entries = append(result.Entries, leaderboard.Entries[offset:offset+limit]...)

	var wg sync.WaitGroup
	wg.Add(len(result.Entries))

	for i, entry := range result.Entries {
		realmSlug := entry.Character.Realm.Slug
		character := entry.Character.Name
		go func(i int) {
			defer wg.Done()
			result.Entries[i].Character.Summary, err = app.client.getCharacterSummary(realmSlug, character, token.AccessToken)
			if err != nil {
				app.serverError(w, err)
				return
			}
		}(i)
	}

	wg.Wait()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func include(entry data.LeaderboardEntry, filters url.Values) bool {
	specs := filters["spec"]
	classes := filters["class"]
	included := false

	if len(classes) > 0 {
		if contains(classes, strings.ToLower(entry.Character.Summary.Class.Name)) {
			included = true
		}

		if contains(specs, strings.ToLower(entry.Character.Summary.Spec.Name)) {
			included = true
		}
	}

	return included
}

func contains(strings []string, value string) bool {
	for _, s := range strings {
		if s == value {
			return true
		}
	}
	return false
}

func (app *application) getCharacterEquipment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	equipment, err := app.client.getEquipmentData(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipment)
}

func (app *application) getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.getToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	summary, err := app.client.getCharacterSummary(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
