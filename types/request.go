package types

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	ResponseWriter http.ResponseWriter
	Route          string
	Method         string
	Values         []byte
}

func (r *Request) SendRequest() ([]byte, error) {

	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Add("Access-Control-Allow-Credentials", "true")

	var reader io.Reader
	if r.Values == nil {
		reader = nil
	} else {
		reader = bytes.NewBuffer(r.Values)
	}

	req, err := http.NewRequest(r.Method, r.Route, reader)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
