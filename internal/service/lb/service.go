package lb

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"testappwink/internal/service"
	"testappwink/pkg/customerr"
)

type srvc struct {
	counter      uint64
	mu           sync.Mutex
	originalHost string
	cdnHost      string
}

type parsedURL struct {
	originalURL     string
	serverClusterID string
	location        string
}

func NewService(originalHost string, cdnHost string) service.LBService {
	return &srvc{
		counter:      0,
		originalHost: originalHost,
		cdnHost:      cdnHost,
	}
}

func (s *srvc) GetURL(ctx context.Context, inputURL string) (string, error) {

	url, err := s.Parse(inputURL)

	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++

	if s.counter%10 == 0 {
		return url.originalURL, nil
	}

	return fmt.Sprintf("http://%s/%s%s", s.cdnHost, url.serverClusterID, url.location), nil
}

func (s *srvc) Parse(inputURL string) (parsedURL, error) {

	url, err := url.ParseRequestURI(inputURL)

	if err != nil {
		return parsedURL{}, customerr.ErrInvalidURL
	}

	host := url.Host

	splittedHost := strings.Split(host, ".")

	if len(splittedHost) < 2 {
		return parsedURL{}, customerr.ErrInvalidURL
	}

	serverClusterID := splittedHost[0]

	location := url.Path

	if len(location) == 0 {
		return parsedURL{}, customerr.ErrInvalidURL
	}

	parsedURL := parsedURL{
		originalURL:     inputURL,
		serverClusterID: serverClusterID,
		location:        location,
	}

	return parsedURL, nil
}
