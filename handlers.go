package main

import (
	"net/http"
)

func (app *application) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
