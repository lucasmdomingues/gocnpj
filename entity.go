package gocnpj

type Company struct {
	MainActivity         []MainActivity        `json:"atividade_principal"`
	SituationDate        string                `json:"data_situacao"`
	Complement           string                `json:"complemento"`
	Type                 string                `json:"tipo"`
	Name                 string                `json:"nome"`
	Email                string                `json:"email"`
	SecondaryActivities  []SecondaryActivities `json:"atividades_secundarias"`
	QSA                  []QSA                 `json:"qsa"`
	Situation            string                `json:"situacao"`
	Neighborhood         string                `json:"bairro"`
	Street               string                `json:"logradouro"`
	Number               string                `json:"numero"`
	ZipCode              string                `json:"cep"`
	City                 string                `json:"municipio"`
	Fantasy              string                `json:"fantasia"`
	Size                 string                `json:"porte"`
	Opening              string                `json:"abertura"`
	LegalNature          string                `json:"natureza_juridica"`
	State                string                `json:"uf"`
	Telephone            string                `json:"telefone"`
	CNPJ                 string                `json:"cnpj"`
	LastUpdate           string                `json:"ultima_atualizacao"`
	Status               string                `json:"status"`
	EFR                  string                `json:"efr"`
	MotivoSituation      string                `json:"motivo_situacao"`
	SituationSpecial     string                `json:"situacao_especial"`
	SituationDateSpecial string                `json:"data_situacao_especial"`
	CapitalSocial        string                `json:"capital_social"`
}

type MainActivity struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type SecondaryActivities struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type QSA struct {
	Qual string `json:"qual"`
	Name string `json:"nome"`
}

type Billing struct {
	Free     bool `json:"free"`
	Database bool `json:"database"`
}

type CompanyError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
