package client

import (
	"encoding/json"
	"formaccount/account"
	"os"
	"reflect"
	"testing"
)

var hostname = getEnv("ACCOUNT_API_HOST", "localhost")
var port = getEnv("ACCOUNT_API_PORT", "8080")

var getEnv = func(name, def string) string {
	envar := os.Getenv(name)
	if envar == "" {
		return def
	}
	return envar
}

var client = NewClient(hostname, port)

func TestList(t *testing.T) {
	assertForSuccessResponse := func(res interface{}, err error) {
		if err != nil {
			t.Fatalf("error returned: %s", err)
		}
		if res == nil {
			t.Error("Account list is empty")
		}
		if reflect.TypeOf(res).String() != "[]*account.Account" {
			t.Errorf("Returned result is not as expected type, have %s", reflect.TypeOf(res).String())
		}
	}

	accounts, err := client.List()
	assertForSuccessResponse(accounts, err)
}

func TestGet(t *testing.T) {
	c := NewClient(hostname, port)

	assertForSuccessResponse := func(res interface{}, err error) {
		if err != nil {
			t.Fatalf("error returned: %s", err)
		}
		if res == nil {
			t.Errorf("Returned result is empty")
		}
		if reflect.TypeOf(res).String() != "*account.Account" {
			t.Errorf("Returned result is not as expected type, have %s", reflect.TypeOf(res).String())
		}
	}

	account, err := c.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assertForSuccessResponse(account, err)
}
func TestCreate(t *testing.T) {
	c := NewClient(hostname, port)

	assertForSuccessResponse := func(res interface{}, err error) {
		if err != nil {
			t.Fatalf("error returned: %s", err)
		}
		if res == nil {
			t.Errorf("Returned result is empty")
		}
		if reflect.TypeOf(res).String() != "*account.Account" {
			t.Errorf("Returned result is not as expected type, have %s", reflect.TypeOf(res).String())
		}
	}

	accountObj := getNewAccount(t, "fd80a4ee-c032-11ea-9af1-23a17806fa35")
	accountCreated, err := c.Create(accountObj)
	t.Logf("Created %v", accountCreated)
	assertForSuccessResponse(accountCreated, err)
}

func TestDelete(t *testing.T) {
	c := NewClient(hostname, port)

	accountObj := getNewAccount(t, "cf10f82a-c032-11ea-9a55-e37fe581f341")
	accountCreated, err := c.Create(accountObj)

	if err != nil {
		t.Fatalf("Could not create account: %s", err)
	}

	t.Log("Account created: " + accountCreated.ID)
	err = c.Delete(accountCreated.ID, accountCreated.Version)

	t.Log("Account deleted: " + accountCreated.ID)
	if err != nil {
		t.Errorf("Account could not be deleted: %s", err)
	}
}

func getNewAccount(t *testing.T, uuid string) account.Account {
	accountJSON, err := getAccountJSON(uuid)
	if err != nil {
		t.Fatal(err)
	}

	var accountObj account.Account
	err = json.Unmarshal([]byte(accountJSON), &accountObj)
	if err != nil {
		t.Fatal(err)
	}

	return accountObj
}

func getAccountJSON(uuid string) (string, error) {
	accountJSON := `{
		"type": "accounts",
		"id": "` + uuid + `",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		"attributes": {
		"country": "PL",
		"base_currency": "CZK",
		"bank_id": "400300",
		"bank_id_code": "GBDSC",
		"bic": "NWBKGB22",
		"alternative_bank_account_names": ["mBank","PKO"]
		}
	  }`
	return accountJSON, nil
}
