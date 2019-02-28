package gocnpj

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CNPJ struct {
	AtividadePrincipal    []*AtividadePrincipal    `json:"atividade_principal"`
	DataSituacao          string                   `json:"data_situacao"`
	Complemento           string                   `json:"complemento"`
	Tipo                  string                   `json:"tipo"`
	Nome                  string                   `json:"nome"`
	Email                 string                   `json:"email"`
	AtividadesSecundarias []*AtividadesSecundarias `json:"atividades_secundarias"`
	QSA                   []*QSA                   `json:"qsa"`
	Situacao              string                   `json:"situacao"`
	Bairro                string                   `json:"bairro"`
	Logradouro            string                   `json:"logradouro"`
	Numero                string                   `json:"numero"`
	Cep                   string                   `json:"cep"`
	Municipio             string                   `json:"municipio"`
	Fantasia              string                   `json:"fantasia"`
	Porte                 string                   `json:"porte"`
	Abertura              string                   `json:"abertura"`
	NaturezaJuridica      string                   `json:"natureza_juridica"`
	UF                    string                   `json:"uf"`
	Telefone              string                   `json:"telefone"`
	CNPJ                  string                   `json:"cnpj"`
	UltimaAtualizacao     string                   `json:"ultima_atualizacao"`
	Status                string                   `json:"status"`
	EFR                   string                   `json:"efr"`
	MotivoSituacao        string                   `json:"motivo_situacao"`
	SituacaoEspecial      string                   `json:"situacao_especial"`
	DataSituacaoEspecial  string                   `json:"data_situacao_especial"`
	CapitalSocial         string                   `json:"capital_social"`
	Billing               *Billing                 `json:"billing"`
}

type AtividadePrincipal struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type AtividadesSecundarias struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type QSA struct {
	Qual string `json:"qual"`
	Nome string `json:"nome"`
}

type Billing struct {
	Free     bool `json:"free"`
	Database bool `json:"database"`
}

func BuscaCNPJ(cnpj string) (*CNPJ, error) {

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
		return nil, fmt.Errorf("Error: %s", resp.Status)
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
