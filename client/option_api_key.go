package client

import "errors"

type APIKeyOption struct {
	APIKey string
}

var _ Option = &APIKeyOption{}

func NewAPIKeyOption(apiKey string) *APIKeyOption {
	return &APIKeyOption{APIKey: apiKey}
}

func (o *APIKeyOption) execute(fClient *FilePursuitClient) error {
	if fClient != nil {
		fClient.apiKey = o.APIKey
		return nil
	}
	return errors.New("file pursuit client cannot be nil upon executing options")
}
