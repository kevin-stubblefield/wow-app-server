package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (app *application) do(req *http.Request, obj interface{}) error {
	resp, err := app.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
