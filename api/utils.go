package api

import (
	gokh_constants "github.com/BRUHItsABunny/gOkHttp/constants"
	"github.com/BRUHItsABunny/go-filepursuit/constants"
	"net/http"
)

func DefaultHeaders(apiKey string) http.Header {
	result := http.Header{
		"user-agent": {constants.UserAgentFlutter},
		"accept":     {gokh_constants.MIMEApplicationJSON},
		"http_test":  {constants.HTTPTest},
	}
	if len(apiKey) > 1 {
		delete(result, "http_test")
		result["x-rapidapi-host"] = []string{constants.RapidAPIHost}
		result["x-rapidapi-key"] = []string{apiKey}
	}
	return result
}
