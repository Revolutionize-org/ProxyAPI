package database

import (
	"sync"

	"gitlab.com/revolutionize1/foward-api/internal/api/database/model"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/postgres"
	"gitlab.com/revolutionize1/foward-api/internal/api/database/redis"
)

func retrieveProxiesFromPostgres(apiKeyProxies []model.ApiKeyProxy, proxyChan chan<- string, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, apiKeyProxy := range apiKeyProxies {
		proxy, err := postgres.RetrieveProxy(apiKeyProxy)
		if err != nil {
			errChan <- err
			return
		}
		proxyChan <- proxy.Format()
	}
}

func storeProxiesInRedis(apiKey *model.ApiKey, proxyChan <-chan string, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for proxy := range proxyChan {
		if err := redis.StoreProxy(apiKey, proxy); err != nil {
			errChan <- err
		}
	}
}

func GetRandomProxyFromApiKey(key string) (string, error) {
	apiKey, err := redis.RetrieveApiKey(key)
	if err != nil {
		return "", err
	}

	randomProxy, err := redis.RetrieveRandomProxy(apiKey)
	if err == nil && randomProxy != "" {
		return randomProxy, nil
	}

	apiKeyProxies, err := postgres.RetrieveApiKeyProxy(apiKey)
	if err != nil {
		return "", err
	}

	proxyChan := make(chan string)
	defer close(proxyChan)

	errChan := make(chan error, 1)
	defer close(errChan)

	var wg sync.WaitGroup
	wg.Add(2)

	go retrieveProxiesFromPostgres(apiKeyProxies, proxyChan, errChan, &wg)
	go storeProxiesInRedis(apiKey, proxyChan, errChan, &wg)

	wg.Wait()

	err = <-errChan
	if err != nil {
		return "", err
	}

	return randomProxy, nil
}
