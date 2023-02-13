package api

import (
	"context"
	"errors"
	"fmt"
	gokh_constants "github.com/BRUHItsABunny/gOkHttp/constants"
	"github.com/BRUHItsABunny/gOkHttp/requests"
	"github.com/BRUHItsABunny/go-filepursuit/constants"
	"net/http"
	"strconv"
)

func SearchRequest(ctx context.Context, apiKey, query string, parameters *Parameters) (*http.Request, error) {
	if parameters == nil {
		return nil, errors.New("parameters can't be nil")
	}

	urlParams := parameters.Encode()
	urlParams["q"] = []string{query}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointSearch,
		requests.NewHeaderOption(DefaultHeaders(apiKey)),
		requests.NewURLParamOption(urlParams),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}

	return req, nil
}

func DiscoverRequest(ctx context.Context, apiKey, link string, parameters *Parameters) (*http.Request, error) {
	if parameters == nil {
		return nil, errors.New("parameters can't be nil")
	}

	urlParams := parameters.Encode()
	urlParams["link"] = []string{link}

	req, err := requests.MakePOSTRequest(ctx, constants.EndpointDiscover,
		requests.NewHeaderOption(DefaultHeaders(apiKey)),
		// I didn't make this up either
		requests.NewURLParamOption(urlParams),
		requests.NewPOSTFormOption(urlParams),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakePOSTRequest: %w", err)
	}

	return req, nil
}

func SuggestRequest(ctx context.Context, apiKey, term string, parameters *Parameters, count int) (*http.Request, error) {
	if parameters == nil {
		return nil, errors.New("parameters can't be nil")
	}

	urlParams := parameters.Encode()
	urlParams["term"] = []string{term}
	urlParams["count"] = []string{strconv.Itoa(count)}

	headers := DefaultHeaders(apiKey)
	headers["content-type"] = []string{gokh_constants.MIMEApplicationJSON}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointSuggest,
		requests.NewHeaderOption(headers),
		requests.NewURLParamOption(urlParams),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}

	return req, nil
}
