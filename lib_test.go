package go_filepursuit

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"github.com/BRUHItsABunny/go-filepursuit/api"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func loadTestEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

var defaultParameters = NewParameters(api.SortTypeDateDESC, api.FileTypeAll, 0)

func TestSearch(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	query := os.Getenv("VAL_SEARCH_QUERY")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c, err := NewFilePursuitClient(hClient)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	res, err := c.Search(context.Background(), query, defaultParameters)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestSuggest(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	term := os.Getenv("VAL_SUGGEST_TERM")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c, err := NewFilePursuitClient(hClient)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	res, err := c.Suggest(context.Background(), term, defaultParameters, 4)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestDiscover(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	link := os.Getenv("VAL_DISCOVER_LINK")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c, err := NewFilePursuitClient(hClient)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	res, err := c.Discover(context.Background(), link, defaultParameters)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}
