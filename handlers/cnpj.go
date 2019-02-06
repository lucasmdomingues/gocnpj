package handlers

import (
	"GoCNPJ/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CnpjHandler(w http.ResponseWriter, r *http.Request) {

	log := types.Log{}

	vars := mux.Vars(r)

	cnpj := vars["cnpj"]
	if cnpj == "" || len(cnpj) < 14 || len(cnpj) > 14 {
		log.Info(w, nil, "CNPJ inválido.")
		return
	}

	c := &types.CNPJ{
		CNPJ: cnpj,
	}

	route := c.MakeURL()

	request := &types.Request{
		ResponseWriter: w,
		Route:          route,
		Method:         "GET",
		Values:         nil,
	}

	response, err := request.SendRequest()
	if err != nil {
		log.Error(w, err)
		return
	}

	err = json.Unmarshal(response, c)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(c)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
