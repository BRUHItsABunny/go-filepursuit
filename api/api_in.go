package api

import (
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/responses"
	"net/http"
)

func InterfaceParser(resp *http.Response) (interface{}, error) {
	result := new(interface{})
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, fmt.Errorf("responses.CheckHTTPCode: %w", err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func ObjectParser[T any](obj T, resp *http.Response) (*T, error) {
	result := new(T)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, fmt.Errorf("responses.CheckHTTPCode: %w", err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}
