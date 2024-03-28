package postgres

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/model"
)

func RetrieveApiKey(inputKey string) (*model.ApiKey, error) {
	var key model.ApiKey

	if err := Instance.Get(&key, "SELECT * FROM api_key WHERE key = $1", inputKey); err != nil {
		if err == sql.ErrNoRows {
			return nil, keyauth.ErrMissingOrMalformedAPIKey
		}
		return nil, err
	}
	return &key, nil
}

func RetrieveApiKeyProxy(apiKey *model.ApiKey) ([]model.ApiKeyProxy, error) {
	var apiKeyProxyStruct []model.ApiKeyProxy

	if err := Instance.Select(&apiKeyProxyStruct, "SELECT * FROM api_key_proxy WHERE api_key_id = $1", apiKey.Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no proxy found")
		}
		return nil, err
	}
	return apiKeyProxyStruct, nil
}

func RetrieveProxy(apiKeyProxy model.ApiKeyProxy) (*model.Proxy, error) {
	var proxy model.Proxy

	if err := Instance.Get(&proxy, "SELECT * FROM proxy WHERE id = $1", apiKeyProxy.ProxyID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("proxy id not found")
		}
		return nil, err
	}
	return &proxy, nil
}
