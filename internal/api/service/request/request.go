package request

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gitlab.com/revolutionize1/foward-api/internal/api/database"
)

const defaultTimeout = 5 * time.Second

func query(req *requestInformation) (*http.Response, error) {
	request, err := createRequest(req)
	if err != nil {
		return nil, err
	}

	return sendRequest(request)
}

func createRequest(requestInfo *requestInformation) (*http.Request, error) {
	method := strings.ToUpper(requestInfo.Method)
	body := strings.NewReader(requestInfo.Body)

	req, err := http.NewRequest(method, requestInfo.URL, body)
	if err != nil {
		return nil, err
	}

	req.Header = parseHeader(requestInfo.Header)
	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	apiKey := req.Header["X-Api-Key"][0]
	proxy, err := database.GetRandomProxyFromApiKey(apiKey)

	if err != nil {
		return nil, err
	}

	log.Debug("Using proxy: " + proxy)

	proxyUrl, err := url.Parse(proxy)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: defaultTimeout,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		closeResponseBody(resp)
		return nil, err
	}

	return resp, nil
}

func closeResponseBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
}
