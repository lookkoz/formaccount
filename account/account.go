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
	AlternativeBankAccountNames []string           `json:"alternative_bank_account_names,omitempty"`
	BankID                      string             `json:"bank_id"`
	BankIDCode                  string             `json:"bank_id_code"`
	BaseCurrency                values.Currency    `json:"base_currency"`
	Bic                         values.Bic         `json:"bic"`
	Country                     values.CountryCode `json:"country"`
	// Iban                        string
	// CustomerId                  string
	// Name                        [4]string
}

/*
 */

/*
type Account struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	OrganisationID string `json:"organisation_id"`
	CreatedOn      time.Time
	ModifiedOn     time.Time
	Version        int               `json:"version"`
	Attributes     AccountAttributes `json:"attributes"`
	//Relationships  AccountRelationships `json:"relationships"`
}

type AccountRelationships struct {
	MasterAccount struct {
		Data []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data"`
	} `json:"master_account"`
	Events AccountEvents `json:"account_events"`
}

type AccountEvents struct {
	Data []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

type AccountAttributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency"`
	AccountNumber           string   `json:"account_number"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	Bic                     string   `json:"bic"`
	Iban                    string   `json:"iban"`
	Name                    []string `json:"name"`
	AlternativeNames        []string `json:"alternative_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Switched                bool     `json:"switched"`
	/*
		PrivateIdentification   struct {
			BirthDate      string `json:"birth_date"`
			BirthCountry   string `json:"birth_country"`
			Identification string `json:"identification"`
			Address        string `json:"address"`
			City           string `json:"city"`
			Country        string `json:"country"`
		} `json:"private_identification"`
		OrganisationIdentification struct {
			Identification string `json:"identification"`
			Actors         []struct {
				Name      []string `json:"name"`
				BirthDate string   `json:"birth_date"`
				Residency string   `json:"residency"`
			} `json:"actors"`
			Address []string `json:"address"`
			City    string   `json:"city"`
			Country string   `json:"country"`
		} `json:"organisation_identification"`
	*
	Status string `json:"status"`
}
*/
