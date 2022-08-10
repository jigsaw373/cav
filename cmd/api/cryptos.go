package main

import (
	"net/http"
	"strings"

	"github.com/jigsaw373/cav/internal/validator"
)

type validationResult struct {
	Crypto  string `json:"crypto"`
	Address string `json:"address"`
	Valid   bool   `json:"valid"`
}

func (app *application) validateAddressHandler(w http.ResponseWriter, r *http.Request) {
	c, err := app.readCryptoParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	c = strings.ToLower(c)

	addr, err := app.readAddressParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	valid, err := validator.ValidateAddress(c, addr)
	if err != nil {
		app.logger.Println(err)
		app.failedValidationResponse(w, r, err)
		return
	}

	res := validationResult{
		Crypto:  c,
		Address: addr,
		Valid:   valid,
	}
	err = app.writeJSON(w, enveloping{"result": res}, http.StatusOK, nil)

	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
