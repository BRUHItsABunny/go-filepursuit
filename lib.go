package go_filepursuit

import (
	"github.com/BRUHItsABunny/go-filepursuit/api"
	"github.com/BRUHItsABunny/go-filepursuit/client"
	"net/http"
)

func NewFilePursuitClient(hClient *http.Client, opts ...client.Option) (*client.FilePursuitClient, error) {
	return client.NewFilePursuitClient(hClient, opts...)
}

func NewParameters(sortType api.SortType, fileType api.FileType, start int) *api.Parameters {
	return api.NewParameters(sortType, fileType, start)
}
