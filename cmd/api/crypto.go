package main

import (
	"net/http"
)

type validationResult struct {
	Crypto  string `json:"crypto"`
	Address string `json:"address"`
	Valid   bool   `json:"valid"`
}

func (app *application) validateAddressHandler(w http.ResponseWriter, r *http.Request) {
	crypto, err := app.readCryptoParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	addr, err := app.readAddressParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	res := validationResult{}

	err = app.writeJSON(w, enveloping{"movie": movie}, http.StatusOK, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
