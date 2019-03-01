package gocnpj

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuscaCNPJ(cnpj string) (*CNPJ, error) {

	cnpj = strings.Replace(cnpj, ".", "", len(cnpj))
	cnpj = strings.Replace(cnpj, "/", "", len(cnpj))
	cnpj = strings.Replace(cnpj, "-", "", len(cnpj))

	url := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%s", cnpj)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %s, Try again.", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c *CNPJ

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	if c.Status == "ERROR" {
		return nil, fmt.Errorf("Error on parse cnpj %s", cnpj)
	}

	return c, err
}
