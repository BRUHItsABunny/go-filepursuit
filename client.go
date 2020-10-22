package go_filepursuit

import (
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
	"strconv"
)

func GetFilePursuitClient() *FilePursuitClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &FilePursuitClient{Client: &httpClient}
}

func (fpc *FilePursuitClient) Search(query, fileType, sort string, page int) ([]FileEntry, error) {
	// sort  (Accepted values: sizeasc, sizedesc, dateasc, datedesc, fileasc, filedesc)
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]FileEntry, 0)
	)

	// ready types arguments
	useCustomTypes := false
	switch fileType {
	case FileTypeArchive:
		break
	case FileTypeMobile:
		break
	case FileTypeEBook:
		break
	case FileTypeAudio:
		break
	case FileTypeVideo:
		break
	case FileTypeAll:
		break
	default:
		useCustomTypes = true
	}
	generalType := fileType
	customType := ""
	if useCustomTypes {
		generalType = FileTypeAll
		customType = fileType
	}

	// ready the page
	if page < 1 {
		page = 1
	}
	page = (page - 1) * 20 // page 1 -> 0, page 2 -> 20, page 10 -> 180

	params := url.Values{
		"q":          []string{query},
		"type":       []string{generalType},
		"sort":       []string{sort},
		"start_page": []string{strconv.Itoa(page)},
	}

	postParams := map[string]string{
		"searchQuery": query,
		"startrow":    strconv.Itoa(page),
		"type":        generalType,
		"fileType":    customType,
		"sort":        sort,
	}

	req, err = fpc.Client.MakePOSTRequestWithParameters(baseURL+endpointSearch, params, postParams, map[string]string{"User-Agent": userAgent})

	if err == nil {
		resp, err = fpc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (fpc *FilePursuitClient) Suggest(term string) ([]string, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]string, 0)
	)

	params := url.Values{
		"term": []string{term},
	}

	postParams := map[string]string{
		"term": term,
	}

	req, err = fpc.Client.MakePOSTRequestWithParameters(baseURL+endpointSuggest, params, postParams, map[string]string{"User-Agent": userAgent})

	if err == nil {
		resp, err = fpc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (fpc *FilePursuitClient) Discover(query string, start int) ([]DiscoveryEntry, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]DiscoveryEntry, 0)
	)

	postParams := map[string]string{
		"searchQuery": query,
		"startrow":    strconv.Itoa(start),
	}

	req, err = fpc.Client.MakePOSTRequestWithParameters(baseURL+endpointDiscover, url.Values{}, postParams, map[string]string{"User-Agent": userAgent})

	if err == nil {
		resp, err = fpc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (fpc *FilePursuitClient) Submit(name, email, file, link string) (string, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]byte, 0)
	)

	postParams := map[string]string{
		"name":  name,
		"email": email,
		"file":  file,
		"link":  link,
	}

	req, err = fpc.Client.MakePOSTRequestWithParameters(baseURL+endpointSubmit, url.Values{}, postParams, map[string]string{"User-Agent": userAgent})

	if err == nil {
		resp, err = fpc.Client.Do(req)
		if err == nil {
			result, err = resp.Bytes()
			return string(result), err
		}
	}
	return "", err
}

func (fpc *FilePursuitClient) News() ([]NewsItem, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]NewsItem, 0)
	)

	params := url.Values{
		"ver":  []string{strconv.Itoa(version)},
		"lang": []string{"en"},
	}

	postParams := map[string]string{}

	req, err = fpc.Client.MakePOSTRequestWithParameters(baseURL+endpointNews, params, postParams, map[string]string{"User-Agent": userAgent})

	if err == nil {
		resp, err = fpc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}
