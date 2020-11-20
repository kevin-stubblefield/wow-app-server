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

func (app *application) showLeaderboard(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	filters := r.URL.Query()
	classes := filters["class"]
	specs := filters["spec"]

	limit, err := strconv.Atoi(filters.Get("limit"))
	if err != nil || limit < 25 {
		limit = 25
	}

	offset, err := strconv.Atoi(filters.Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	leaderboard, err := app.leaderboard.FetchAllByBracket(pvpBracket, classes, specs, limit, offset)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "leaderboard.page.tmpl", &templateData{
		Leaderboard: leaderboard,
	})
}

func (app *application) getLeaderboardByBracket(w http.ResponseWriter, r *http.Request) {
	pvpBracket := mux.Vars(r)["bracket"]

	filters := r.URL.Query()
	classes := filters["class"]
	specs := filters["spec"]

	limit, err := strconv.Atoi(filters.Get("limit"))
	if err != nil || limit < 25 {
		limit = 25
	}

	offset, err := strconv.Atoi(filters.Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	leaderboard, err := app.leaderboard.FetchAllByBracket(pvpBracket, classes, specs, limit, offset)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func (app *application) getCharacter(w http.ResponseWriter, r *http.Request) {
	realmSlug := mux.Vars(r)["realmSlug"]
	characterName := mux.Vars(r)["character"]

	character, err := app.character.Fetch(realmSlug, characterName)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}

func (app *application) getSpecs(w http.ResponseWriter, r *http.Request) {
	specs, err := app.specs.FetchSpecs()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specs)
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

	token, err := app.client.GetToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	pvpSeason, err := app.client.GetCurrentPvPSeason(token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	leaderboard, err := app.client.GetLeaderboardData(pvpSeason.CurrentSeason.ID, pvpBracket, token.AccessToken)
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
			summary, err := app.client.GetCharacterSummary(realmSlug, character, token.AccessToken)
			if err != nil {
				app.serverError(w, err)
				return
			}
			result.Entries[i].Character.Summary = *summary
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

func (app *application) getCharacterEquipmentFromBlizzard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.GetToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	equipment, err := app.client.GetCharacterEquipment(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipment)
}

func (app *application) getCharacterFromBlizzard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	realmSlug := vars["realmSlug"]
	character := vars["character"]

	token, err := app.client.GetToken()
	if err != nil {
		app.serverError(w, err)
		return
	}

	summary, err := app.client.GetCharacterSummary(realmSlug, character, token.AccessToken)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
