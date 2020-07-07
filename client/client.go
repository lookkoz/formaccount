package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"formaccount/account"
	"io"
	"net/http"
	"time"
)

// var ErrNoRows = errors.New("form3.client: no rows in result set")
// var ErrNotFound = errors.New("form3.client: not found")
// var ErrConn = errors.New("form3.client: rest service is down")

const (
	ClientName    = "form3-account-client"
	ClientVersion = "v1"

	DefaultHost = "http://localhost:8080"
	BasePath    = "/v1/organisation/accounts"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	UserAgent  string
}

func NewClient(accountServiceEndpoint string) *Client {
	tr := &http.Transport{
		MaxIdleConns:       5,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	return &Client{
		httpClient: client,
		UserAgent:  ClientName,
		baseURL:    accountServiceEndpoint,
	}
}

// List returns list of accounts
func (c *Client) List(page account.Page) ([]*account.Account, error) {
	reqURL := fmt.Sprintf("%s/v1/organisation/accounts%s", c.baseURL, page.String())

	req, err := c.getRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accountsResponse account.AccountsResponse
	err = json.NewDecoder(resp.Body).Decode(&accountsResponse)
	return accountsResponse.Accounts, err
}

// Fetch returns one account
func (c Client) Fetch(id string) (*account.Account, error) {
	reqURL := fmt.Sprintf("%s/v1/organisation/accounts/%s", c.baseURL, id)

	req, err := c.getRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accountResponse account.AccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	return accountResponse.AccountObject, err
}

func (c Client) Create(accountObj account.Account) (*account.Account, error) {
	reqURL := fmt.Sprintf("%s/v1/organisation/accounts", c.baseURL)

	requestBody := account.AccountRequest{Account: &accountObj}
	req, err := c.getRequest(http.MethodPost, reqURL, requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req, 201)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accountResponse account.AccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	return accountResponse.AccountObject, err
}

func (c Client) Delete(id string, version int64) error {
	reqURL := fmt.Sprintf(c.baseURL+"/v1/organisation/accounts/%s?version=%d", id, version)
	req, err := c.getRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req, 204)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) getRequest(method, url string, requestBody interface{}) (*http.Request, error) {
	body, err := getJsonRequestBody(requestBody)
	req, err := http.NewRequest(method, url, body)

	if body != nil {
		req.Header.Set("Content-Type", "application/api+json")
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/api+json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) doRequest(req *http.Request, expectedStatus int) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectedStatus {
		var cErr ClientError
		err = json.NewDecoder(resp.Body).Decode(&cErr)
		cErr.HttpStatusCode = resp.StatusCode
		return nil, cErr
	}

	return resp, nil
}

func getJsonRequestBody(requestBody interface{}) (io.Reader, error) {
	var body io.Reader
	if requestBody != nil {
		buffer, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("JSON Enc error: %s", err)
		}
		body = bytes.NewBuffer(buffer)
	}
	return body, nil
}
