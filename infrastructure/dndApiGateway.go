package infrastructure

import (
	"context"
	"fmt"
	"io"
	"net/http"
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
		rateLimiter: rate.NewLimiter(rate.Limit(10), 10), // 10 tokens/sec, max tokens is 10
	}
}

func (dndApiGateway *DndApiGateway) Get(endpoint string) ([]byte, error) {
	err := dndApiGateway.rateLimiter.Wait(context.Background())
	if !(err == nil) {
		return nil, err
	}

	resp, err := dndApiGateway.client.Get(dndApiGateway.baseUrl + endpoint)
	if !(err == nil) {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-200 response: %d", resp.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if !(err == nil) {
		return nil, err
	}

	return body, nil
}
