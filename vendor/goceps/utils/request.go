package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func MakeRequest(w http.ResponseWriter, method, route string, jsonValue []byte) ([]byte, error) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	req, err := http.NewRequest(method, route, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}
