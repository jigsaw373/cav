package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := enveloping{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	if err := app.writeJSON(w, env, http.StatusOK, nil); err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
