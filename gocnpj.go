package gocnpj

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
)

type service struct {
	BaseURL url.URL
}

type Service interface {
	Search(cnpj string) (*Company, error)
}

func NewService() Service {
	return &service{
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "www.receitaws.com.br",
			Path:   "v1/",
		},
	}
}

func (s *service) Search(cnpj string) (*Company, error) {
	exp := regexp.MustCompile(`[^0-9]`)
	cnpj = exp.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return nil, errors.New("cnpj must be 14 characters")
	}

	url, err := s.BaseURL.Parse(path.Join("cnpj", cnpj))
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var companyError *CompanyError

	err = json.Unmarshal(body, &companyError)
	if err != nil {
		return nil, err
	}

	if companyError.Status != "OK" {
		return nil, errors.New(companyError.Message)
	}

	var company *Company

	err = json.Unmarshal(body, &company)
	if err != nil {
		return nil, err
	}

	return company, err
}
