package account

type Response struct {
	Status       int
	ErrorMessage string `json:"error_message,omitempty"`
}

type AccountResponse struct {
	Response
	Account *Account `json:"data"`
	Links   struct {
		Self string `json:"data"`
	} `json:"links"`
}

type AccountsResponse struct {
	Response
	Accounts []*Account `json:"data"`
	Links    struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	} `json:"links"`
}
