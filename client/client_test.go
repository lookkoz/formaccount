package client

import (
	"encoding/json"
	"formaccount/account"
	"os"
	"reflect"
	"testing"
)

var endpoint = getEnv("ACCOUNT_API_ENDPOINT", "http://localhost:8080")

var getEnv = func(name, def string) string {
	envar := os.Getenv(name)
	if envar == "" {
		return def
	}
	return envar
}

var accountService = NewClient(endpoint)

func init() {
	accountService.Delete("cf10f82a-c032-11ea-9a55-e37fe581f341", 0)
	accountService.Delete("fd80a4ee-c032-11ea-9af1-23a17806fa35", 0)
	accountService.Delete("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", 0)
}

/** POSITVIE SCENARIOS **/

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

	accounts, err := accountService.List(account.Page{})
	assertForSuccessResponse(accounts, err)
}

func TestListPagination(t *testing.T) {
	// create few accounts for tests first
	ids := []string{
		"1180a4ee-c032-11ea-9af1-23a17806fa35",
		"2280a4ee-c032-11ea-9af1-23a17806fa35",
		"3380a4ee-c032-11ea-9af1-23a17806fa35",
	}

	createAccounts(ids, t)

	pageSize := 2
	page := account.Page{}
	page.Size(pageSize).Number(1)
	accounts, err := accountService.List(page)

	if err != nil {
		t.Errorf("Not expected error: %s", err.Error())
	}

	if len(accounts) != pageSize {
		t.Errorf("Expected %d elements, got %d", pageSize, len(accounts))
	}

	page.Number(100).Size(10)
	accounts, err = accountService.List(page)
	if len(accounts) != 0 {
		t.Errorf("Expected %d elements, got %d", pageSize, len(accounts))
	}

	removeAccounts(ids, t)
}

func TestGet(t *testing.T) {
	accountService.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)

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

	uuid := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	accountObj := getNewAccount(t, uuid)
	_, err := accountService.Create(accountObj)
	if err != nil {
		t.Fatalf("Could not create account: %s", err)
	}
	account, err := accountService.Fetch(uuid)
	assertForSuccessResponse(account, err)
}
func TestCreate(t *testing.T) {
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
	accountCreated, err := accountService.Create(accountObj)
	t.Logf("Created %v", accountCreated)
	assertForSuccessResponse(accountCreated, err)
}

func TestDelete(t *testing.T) {
	accountObj := getNewAccount(t, "cf10f82a-c032-11ea-9a55-e37fe581f341")
	accountCreated, err := accountService.Create(accountObj)

	if err != nil {
		t.Fatalf("Could not create account: %s", err)
	}

	t.Log("Account created: " + accountCreated.ID)
	err = accountService.Delete(accountCreated.ID, accountCreated.Version)

	t.Log("Account deleted: " + accountCreated.ID)
	if err != nil {
		t.Errorf("Account could not be deleted: %s", err)
	}
}

/** NEGATIVE SCENARIOS **/

func TestCreateValidationError(t *testing.T) {
	accountObj := getNewAccount(t, "fd80a4ee-c032-11ea-9af1-23a17806fa35")
	accountObj.Attributes.Country = ""
	accountCreated, err := accountService.Create(accountObj)

	assertForFailureResponse(t, accountCreated, err)
}

func TestFetchNotExistingError(t *testing.T) {
	account, err := accountService.Fetch("fd80a4ee-c032-11ea-0000-23a17806fa00")
	assertForFailureResponse(t, account, err)
}

func TestDeleteValidationError(t *testing.T) {
	err := accountService.Delete("AOS00000000-c032-11ea-0000-23a17806fa00", 1000)
	assertForFailureResponse(t, nil, err)
}

func TestWrongEndpointError(t *testing.T) {
	accountService.baseURL = "http//localhost:8080"
	_, err := accountService.Fetch("fd80a4ee-c032-11ea-0000-23a17806fa00")
	if err == nil {
		t.Error("Expected error")
	}
}

func assertForFailureResponse(t *testing.T, res *account.Account, err error) {
	t.Helper()

	if err == nil {
		t.Error("Expected error")
	}

	if res != nil {
		t.Errorf("Not expected value of %T", res)
	}

	if reflect.TypeOf(err).String() != reflect.TypeOf(ClientError{}).String() {
		t.Errorf("Expected %T , got %T", ClientError{}, err)
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

func createAccounts(ids []string, t *testing.T) {
	for _, id := range ids {
		_, err := accountService.Create(getNewAccount(t, id))
		if err != nil {
			t.Fatal("Cannot create account: ", err)
		}
	}
}

func removeAccounts(ids []string, t *testing.T) {
	for _, id := range ids {
		err := accountService.Delete(id, 0)
		if err != nil {
			t.Fatal("Cannot remove account: ", err)
		}
	}
}
