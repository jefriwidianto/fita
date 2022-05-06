package Controller

import (
	"encoding/json"
	"fita/Controller/DTO/request"
	"fita/Helper"
	"fita/Repository/transaction"
	"net/http"
)

type CtrlTransaction struct {}

type Transaction interface {
	Transaction(w http.ResponseWriter, r *http.Request)
}

func (c CtrlTransaction) Transaction(w http.ResponseWriter, r *http.Request) {
	var params request.ReqBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Helper.HttpResponseError(w, r, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data, err := transaction.RepositoryNew(params).SchemaItems("Items").Result()
	if err != nil {
		Helper.HttpResponseError(w, r, err, http.StatusInternalServerError)
		return
	}

	Helper.HttpResponseSuccess(w, r, data.Data)
}