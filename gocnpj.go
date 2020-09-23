package gocnpj

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type service struct{}

type Service interface {
	Search(cnpj string) (*Company, error)
}

func NewService() Service {
	return new(service)
}

func (s *service) Search(cnpj string) (*Company, error) {
	cnpj = strings.Replace(cnpj, ".", "", len(cnpj))
	cnpj = strings.Replace(cnpj, "/", "", len(cnpj))
	cnpj = strings.Replace(cnpj, "-", "", len(cnpj))

	url := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%s", cnpj)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Oops, error on search company")
	}

	var companyError *CompanyError

	err = json.Unmarshal(body, &companyError)
	if err != nil {
		return nil, err
	}

	if companyError.Status != "OK" {
		return nil, fmt.Errorf("Oops, %s", companyError.Message)
	}

	var company *Company

	err = json.Unmarshal(body, &company)
	if err != nil {
		return nil, err
	}

	return company, err
}
