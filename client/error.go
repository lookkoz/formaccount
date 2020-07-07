package client

import "fmt"

type ClientError struct {
	HttpStatusCode int
	Message        string `json:"error_message"`
}

func (x ClientError) Error() string {
	return fmt.Sprintf("%s ERR: %s (%d)", ClientName, x.Message, x.HttpStatusCode)
}
