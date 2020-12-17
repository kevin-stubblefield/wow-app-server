package main

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"stubblefield.io/wow-leaderboard-api/models"
)

type templateData struct {
	CSRFToken   string
	CurrentYear int
	Flash       string
	Limit       int
	Offset      int
	Leaderboard []models.LeaderboardEntry
	Character   models.Character
	Slots       []string

	Breakdown2s   []models.ClassAndSpecBreakdown
	Breakdown3s   []models.ClassAndSpecBreakdown
	BreakdownRBGs []models.ClassAndSpecBreakdown
}

func classSlug(class string) string {
	strs := strings.Split(strings.ToLower(class), " ")
	return strings.Join(strs, "-")
}

func wowheadLink(e models.Equipment) string {
	return fmt.Sprintf("https://www.wowhead.com/item=%d?bonus=%s", e.ItemID, e.Bonuses)
}

func findEquipmentForSlot(equipped []models.Equipment, slot string) models.Equipment {
	for _, item := range equipped {
		if item.ItemSlot == slot {
			return item
		}
	}
	return models.Equipment{}
}

func formatFloat(f float32, precision int) string {
	return fmt.Sprintf("%.*f", precision, f)
}

func previousPage(limit, offset int) int {
	result := offset - limit - limit
	if result < 0 {
		result = 0
	}
	return result
}

var functions = template.FuncMap{
	"classSlug":            classSlug,
	"wowheadLink":          wowheadLink,
	"findEquipmentForSlot": findEquipmentForSlot,
	"formatFloat":          formatFloat,
	"previousPage":         previousPage,
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
