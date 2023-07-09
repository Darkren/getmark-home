package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// httpService implements Service by requesting outer authorization service via HTTP.
type httpService struct {
	baseURL string
	httpCl  *http.Client
}

// NewHTTPService constructs new httpService.
func NewHTTPService(httpCl *http.Client, urlStr string) (Service, error) {
	_, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %w", err)
	}

	return &httpService{
		baseURL: urlStr,
		httpCl:  httpCl,
	}, nil
}

// Auth authorizes user with the provided login and returns issued JWT token.
func (s *httpService) Auth(login string) (string, error) {
	if err := s.ping(); err != nil {
		return "", fmt.Errorf("ping: %w", err)
	}

	token, err := s.doRequest(http.MethodGet, s.baseURL+"/generate?login="+login, nil)
	if err != nil {
		return "", fmt.Errorf("doRequest: %w", err)
	}

	return string(token), nil
}

// ValidateToken validates provided token.
func (s *httpService) ValidateToken(token string) error {
	if err := s.ping(); err != nil {
		return fmt.Errorf("ping: %w", err)
	}

	_, err := s.doRequest(http.MethodGet, s.baseURL+"/validate", map[string]string{
		"Authorization": "Bearer " + token,
	})
	if err != nil {
		return fmt.Errorf("doRequest: %w", err)
	}

	return nil
}

// ping tests authorization service availability.
func (s *httpService) ping() error {
	respBody, err := s.doRequest(http.MethodGet, s.baseURL+"/ping", nil)
	if err != nil {
		return fmt.Errorf("doRequest: %w", err)
	}

	if string(respBody) != "pong" {
		return fmt.Errorf("invalid response: %s", respBody)
	}

	return nil
}

// doRequest performs HTTP request and returns response body.
func (s *httpService) doRequest(method, url string, headers map[string]string) ([]byte, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	for headerName, headerVal := range headers {
		request.Header.Add(headerName, headerVal)
	}

	resp, err := s.httpCl.Do(request)
	if err != nil {
		return nil, fmt.Errorf("httpCl.Get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return respBody, nil
}
