package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"formaccount/account"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// var ErrNoRows = errors.New("form3.client: no rows in result set")
// var ErrNotFound = errors.New("form3.client: not found")
// var ErrConn = errors.New("form3.client: rest service is down")

type ClientError struct {
	HttpStatusCode int
	Message        string `json:"error_message"`
}

func (x ClientError) Error() string {
	return fmt.Sprintf("%s ERR: %s (%d)", ClientName, x.Message, x.HttpStatusCode)
}

const (
	ClientName    = "form3-account-client"
	ClientVersion = "v1"

	DefaultHostname = "localhost"
	DefaultPort     = "8080"
	UserAgentName   = ClientName
	DefaultScheme   = "http"
	BasePath        = "/v1/organisation/accounts"
)

type Client struct {
	httpClient *http.Client
	BaseURL    *url.URL
	UserAgent  string
}

func NewClient(hostname, port string) *Client {
	tr := &http.Transport{
		MaxIdleConns:       5,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	url := getURL(hostname, port)

	return &Client{
		httpClient: client,
		UserAgent:  UserAgentName,
		BaseURL:    url,
	}
}

func getURL(hostname, port string) *url.URL {
	if port != "" {
		port = ":" + port
	}

	url := &url.URL{
		Host:   hostname + port,
		Scheme: DefaultScheme,
		Path:   BasePath,
	}

	return url
}

// List returns list of accounts
func (c *Client) List() ([]*account.Account, error) {
	c.BaseURL.Path = "/v1/organisation/accounts"
	u := c.BaseURL.ResolveReference(c.BaseURL)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
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
	c.BaseURL.Path = "/v1/organisation/accounts/" + id
	req, err := http.NewRequest(http.MethodGet, c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var cErr ClientError
		err = json.NewDecoder(resp.Body).Decode(&cErr)
		cErr.HttpStatusCode = resp.StatusCode
		return nil, cErr
	}

	var accountResponse account.AccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	return accountResponse.Account, err
}

func (c Client) Create(accountObj account.Account) (*account.Account, error) {
	c.BaseURL.Path = "/v1/organisation/accounts/"

	requestBody := account.AccountRequest{Account: &accountObj}
	req, err := c.getRequest(http.MethodPost, c.BaseURL.String(), requestBody)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var response account.Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		return nil, errors.New(response.ErrorMessage)
	}

	var accountResponse account.AccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	return accountResponse.Account, err
}

func (c Client) Delete(id string, version int64) error {
	c.BaseURL.Path = "/v1/organisation/accounts"
	url := c.BaseURL
	url.Query().Set("version", strconv.Itoa(int(version)))

	req, err := http.NewRequest(http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil
	}

	return nil
}

func (c *Client) getRequest(method, url string, requestBody interface{}) (*http.Request, error) {
	var body io.Reader
	if requestBody != nil {
		buffer, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("JSON encoding error: %s", err)
		}
		body = bytes.NewBuffer(buffer)
	}
	req, err := http.NewRequest(method, url, body)

	if body != nil {
		req.Header.Set("Content-Type", "application/api+json")
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/api+json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, err
}
