package account

import (
	"formaccount/account/values"
	"time"
)

type Account struct {
	ID             string            `json:"id"`
	Type           string            `json:"type"`
	OrganisationID string            `json:"organisation_id"`
	CreatedOn      time.Time         `json:"created_on"`
	ModifiedOn     time.Time         `json:"modified_on"`
	Version        int64             `json:"version"`
	Attributes     AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	AlternativeBankAccountNames []string                   `json:"alternative_bank_account_names,omitempty"`
	BankID                      string                     `json:"bank_id,omitempty"`
	BankIDCode                  string                     `json:"bank_id_code,omitempty"`
	BaseCurrency                values.Currency            `json:"base_currency,omitempty"`
	Bic                         values.Bic                 `json:"bic,omitempty"`
	Country                     values.CountryCode         `json:"country"`
	Iban                        string                     `json:"iban,omitempty"`
	CustomerId                  string                     `json:"customer_id,omitempty"`
	Name                        [4]string                  `json:"name"`
	AlternativeNames            [3]string                  `json:"alternative_names,omitempty"`
	AccountClassification       string                     `json:"account_classification,omitempty"`
	JointAccount                bool                       `json:"joint_account,omitempty"`
	AccountMatchingOptOut       bool                       `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification     string                     `json:"secondary_identification,omitempty"`
	Switched                    bool                       `json:"switched,omitempty"`
	PrivateIdentification       PrivateIdentification      `json:"private_identification"`
	OrganisationIdentification  OrganisationIdentification `json:"organisation_identification"`
}

type PrivateIdentification struct {
	BirthDate      string   `json:"birth_date"`
	BirthCountry   string   `json:"birth_country"`
	Identification string   `json:"identification"`
	Address        []string `json:"address"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
}

type OrganisationIdentification struct {
	Identification string   `json:"identification"`
	Actors         Actors   `json:"actors"`
	Address        []string `json:"address"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
}

type Actors []struct {
	Name      []string `json:"name"`
	BirthDate string   `json:"birth_date"`
	Residency string   `json:"residency"`
}
