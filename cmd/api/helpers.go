package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type enveloping map[string]interface{}

func (app *application) readCryptoParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	crypto := params.ByName("crypto")
	if crypto == "" {
		return "", errors.New("invalid crypto")
	}

	return crypto, nil
}

func (app *application) readAddressParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	addr := params.ByName("address")
	if addr == "" {
		return "", errors.New("invalid address")
	}

	return addr, nil
}

func (app *application) writeJSON(w http.ResponseWriter, data interface{}, statusCode int, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
	return nil
}

func (app *application) readJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contain badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contain badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contain incorrect json fields for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contain incorrect json fields (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}

	}

	// TODO: Handle all error types for reading data

	return nil
}
