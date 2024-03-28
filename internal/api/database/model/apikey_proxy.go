package model

type ApiKeyProxy struct {
	ApiKeyID string `db:"api_key_id"`
	ProxyID  string `db:"proxy_id"`
}
