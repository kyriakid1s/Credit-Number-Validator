package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"creditvalidator.kortsi.org/internal/validator"
	luhnalgor "creditvalidator.kortsi.org/microservices/luhnAlgor"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("paok"))
}

func (app *application) validateCreditCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")

	validator := validator.Validator{}
	params := httprouter.ParamsFromContext(r.Context())
	cardNumber := params.ByName("number")

	validator.CheckField(validator.NotBlank(cardNumber), "number", "Card Number can't be blank!")
	validator.CheckField(validator.IsNumeric(cardNumber), "number", "Card Number can only contains Integers!")
	validator.CheckField(validator.CorrectLength(cardNumber), "number", "Card Number must be 16 digits!")

	if !validator.Valid() {
		res, err := json.Marshal(validator.FieldErrors)
		if err != nil {

			app.serverError(w, r, err)
			return
		}
		http.Error(w, string(res), http.StatusUnprocessableEntity)
		return
	}

	isValid, err := luhnalgor.Validate(cardNumber)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.Write([]byte(strconv.FormatBool(isValid)))

}
