package account

import "fmt"

type AccountRequest struct {
	Account *Account `json:"data"`
}

type Page struct {
	size   int
	number int
}

func (p *Page) Size(s int) *Page {
	if s > 0 {
		p.size = s
	}
	return p
}

func (p *Page) Number(n int) *Page {
	if n > 0 {
		p.number = n
	}
	return p
}

func (p *Page) String() string {
	return fmt.Sprintf("?page[number]=%d&page[size]=%d", p.number, p.size)
}
