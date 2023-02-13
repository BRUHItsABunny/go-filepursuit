package client

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/go-filepursuit/api"
	"net/http"
)

type FilePursuitClient struct {
	Client *http.Client
	apiKey string
}

func NewFilePursuitClient(hClient *http.Client, opts ...Option) (*FilePursuitClient, error) {
	if hClient == nil {
		hClient = http.DefaultClient
	}
	result := &FilePursuitClient{
		Client: hClient,
		apiKey: "",
	}

	for optIdx, opt := range opts {
		err := opt.execute(result)
		if err != nil {
			return nil, fmt.Errorf("opt.execute[%d]: %w", optIdx, err)
		}
	}
	return result, nil
}

func (c *FilePursuitClient) Search(ctx context.Context, query string, parameters *api.Parameters) ([]*api.FPEntry, error) {
	req, err := api.SearchRequest(ctx, c.apiKey, query, parameters)
	if err != nil {
		return nil, fmt.Errorf("api.SearchRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.FPResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result.FilesFound, nil
}

func (c *FilePursuitClient) Suggest(ctx context.Context, term string, parameters *api.Parameters, count int) ([]string, error) {
	req, err := api.SuggestRequest(ctx, c.apiKey, term, parameters, count)
	if err != nil {
		return nil, fmt.Errorf("api.SuggestRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser([]string{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return *result, nil
}

func (c *FilePursuitClient) Discover(ctx context.Context, link string, parameters *api.Parameters) ([]*api.FPEntry, error) {
	req, err := api.DiscoverRequest(ctx, c.apiKey, link, parameters)
	if err != nil {
		return nil, fmt.Errorf("api.DiscoverRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.FPResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result.Result, nil
}
