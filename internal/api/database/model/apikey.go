package model

import "time"

type ApiKey struct {
	Id             string
	Key            string
	IpAddress      string    `db:"ip_address"`
	NumProxies     int       `db:"num_proxies"`
	CreatedAt      time.Time `db:"created_at"`
	ExpirationDate time.Time `db:"expiration_date"`
}
