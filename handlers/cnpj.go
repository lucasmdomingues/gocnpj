package handlers

import (
	"encoding/json"
	"fmt"
	"gocnpj/types"
	"gocnpj/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var l utils.Log

func CnpjHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	cnpj := vars["cnpj"]

	if cnpj == "" || len(cnpj) < 8 {
		l.Info(w, nil, "CNPJ is required")
		return
	}

	route := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%v", cnpj)

	data, err := utils.MakeRequest(w, "GET", route, nil)
	if err != nil {
		l.Info(w, nil, err.Error())
	}

	var c types.CNPJ

	err = json.Unmarshal(data, &c)
	if err != nil {
		l.Info(w, nil, err.Error())
	}

	l.Info(w, c, "request success")
}
