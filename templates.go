package main

import (
	"html/template"
	"path/filepath"
	"strings"

	"stubblefield.io/wow-leaderboard-api/models"
)

type templateData struct {
	CSRFToken   string
	CurrentYear int
	Flash       string
	Leaderboard []models.LeaderboardEntry
	Character   []models.Character
}

func classSlug(class string) string {
	strs := strings.Split(strings.ToLower(class), " ")
	return strings.Join(strs, "-")
}

var functions = template.FuncMap{
	"classSlug": classSlug,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
