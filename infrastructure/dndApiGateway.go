package infrastructure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type DndApiGateway struct {
	baseUrl     string
	client      *http.Client
	rateLimiter *rate.Limiter
}

func NewDndApiGateway(baseUrl string) *DndApiGateway {
	return &DndApiGateway{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		rateLimiter: rate.NewLimiter(rate.Limit(10), 10),
	}
}

func (dndApiGateway *DndApiGateway) Get(endpoint string) ([]byte, error) {
	err := dndApiGateway.rateLimiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}

	response, err := dndApiGateway.client.Get(dndApiGateway.baseUrl + endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 response: %d", response.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (dndApiGateway *DndApiGateway) GetMultipleOrdered(endpoints []string) ([][]byte, []error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	bodies := make([][]byte, len(endpoints))
	errors := []error{}

	for i, endpoint := range endpoints {
		wg.Add(1)

		go func(i int, endpoint string) {
			defer wg.Done()

			body, err := dndApiGateway.Get(endpoint)
			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("request failed for %s: %w", endpoint, err))
				mu.Unlock()
				return
			}

			bodies[i] = body
		}(i, endpoint)
	}

	wg.Wait()
	return bodies, errors
}

func (dndApiGateway *DndApiGateway) GetMultipleUnordered(endpoints []string) ([][]byte, []error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	bodies := make([][]byte, 0, len(endpoints))
	errors := []error{}

	for _, endpoint := range endpoints {
		wg.Add(1)

		go func(endpoint string) {
			defer wg.Done()

			body, err := dndApiGateway.Get(endpoint)
			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errors = append(errors, fmt.Errorf("request failed for %s: %w", endpoint, err))
			} else {
				bodies = append(bodies, body)
			}
		}(endpoint)
	}

	wg.Wait()
	return bodies, errors
}
