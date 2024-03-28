package model

import "fmt"

type Proxy struct {
	Id       string
	Address  string
	Username string
	Password string
	Scheme   string
}

func (p *Proxy) Format() string {
	return fmt.Sprintf("%s://%s:%s@%s", p.Scheme, p.Username, p.Password, p.Address)
}
