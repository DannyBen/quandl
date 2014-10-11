package quandl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ApiKey is used to set your api key before you make any call
var ApiKey = ""

// LastUrl will hold the last requested URL after each call
var LastUrl = ""

// CacheHandler is a reference to a struct that implements the Cacher interface.
// If set, it will use it to get documents from the cache or set to it.
// See the example cache handler in the quandl/cache package.
var CacheHandler Cacher

var urlTemplates map[string]string = map[string]string{
	"symbol":  "https://www.quandl.com/api/v1/datasets/%s.%s?%s",
	"symbols": "https://www.quandl.com/api/v1/multisets.%s?columns=%s&%s",
	"search":  "https://www.quandl.com/api/v1/datasets.%s?%s",
	"list":    "https://www.quandl.com/api/v2/datasets.%s?%s",
	// "favs":    "https://www.quandl.com/api/v1/current_user/collections/datasets/favourites.%s?auth_token=%s",
}

// Type Options is used to send additional parameters in the Quandl request
type Options url.Values

// Set registers a key=value pair to be sent in the Quandl request
func (o Options) Set(key, value string) {
	o[key] = []string{value}
}

type Cacher interface {
	Get(key string) []byte
	Set(key string, data []byte) error
}

// GetSymbol returns data for a given symbol
func GetSymbol(symbol string, params Options) (*SymbolResponse, error) {
	raw, err := GetSymbolRaw(symbol, "json", params)
	var response SymbolResponse
	if err != nil {
		return &response, err
	}

	err = json.Unmarshal(raw, &response)
	if err != nil {
		return &response, marshallerError(raw, err)
	}
	return &response, nil
}

// GetSymbols returns data for a symbols array
func GetSymbols(symbols []string, params Options) (*SymbolsResponse, error) {
	raw, err := GetSymbolsRaw(symbols, "json", params)
	var response SymbolsResponse
	if err != nil {
		return &response, err
	}

	err = json.Unmarshal(raw, &response)
	if err != nil {
		return &response, marshallerError(raw, err)
	}
	return &response, nil
}

// GetList returns a list of symbols for a source
func GetList(source string, page int, perPage int) (*ListResponse, error) {
	raw, err := GetListRaw(source, "json", page, perPage)
	var response ListResponse
	if err != nil {
		return &response, err
	}

	err = json.Unmarshal(raw, &response)
	if err != nil {
		return &response, marshallerError(raw, err)
	}
	return &response, nil
}

// GetSearch returns search results
func GetSearch(query string, page int, perPage int) (*SearchResponse, error) {
	raw, err := GetSearchRaw(query, "json", page, perPage)
	var response SearchResponse
	if err != nil {
		return &response, err
	}

	err = json.Unmarshal(raw, &response)
	if err != nil {
		return &response, marshallerError(raw, err)
	}
	return &response, nil
}

// GetSymbolRaw returns CSV, JSON or XML data for a given symbol
func GetSymbolRaw(symbol string, format string, params Options) ([]byte, error) {
	url := getUrl("symbol", symbol, format, arrangeParams(params))
	return getData(url)
}

// GetSymbolsRaw returns CSV, JSON or XML data for multiple symbols
func GetSymbolsRaw(symbols []string, format string, params Options) ([]byte, error) {
	url := getUrl("symbols", format, symbolsToString(symbols), arrangeParams(params))
	return getData(url)
}

// GetListRaw returns a list of symbols for a source as CSV, JSON or XML
func GetListRaw(source string, format string, page int, perPage int) ([]byte, error) {
	params := Options{}

	params.Set("query", "*")
	params.Set("source_code", source)
	params.Set("per_page", strconv.Itoa(perPage))
	params.Set("page", strconv.Itoa(page))

	url := getUrl("list", format, arrangeParams(params))
	return getData(url)
}

// GetSearchRaw returns search results as JSON or XML
func GetSearchRaw(query string, format string, page int, perPage int) ([]byte, error) {
	params := Options{}

	// TODO: Remove when Quandl fixes this bug
	if format == "csv" {
		format = "json"
	}

	params.Set("query", query)
	params.Set("per_page", strconv.Itoa(perPage))
	params.Set("page", strconv.Itoa(page))

	url := getUrl("search", format, arrangeParams(params))
	return getData(url)
}

// getData requests a URL from Quandl and returns the raw response string
func getData(url string) ([]byte, error) {
	if CacheHandler != nil {
		if response := CacheHandler.Get(url); response != nil {
			return response, nil
		}
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if CacheHandler != nil {
		if err := CacheHandler.Set(url, contents); err != nil {
			return nil, err
		}
	}

	return contents, nil
}

// getUrl receives a kind that points to a URL template and
// a variable number of strings, which will be replaced
// in the template.
func getUrl(kind string, args ...interface{}) string {
	template := urlTemplates[kind]
	LastUrl = strings.Trim(fmt.Sprintf(template, args...), "&?")
	return LastUrl
}

// arrangeParams takes an Options map and converts it to
// a query string. It will also append the api key as needed.
func arrangeParams(qs Options) string {
	if qs == nil {
		if ApiKey == "" {
			return ""
		}
		qs = Options{}
	}
	if ApiKey != "" {
		qs.Set("auth_token", ApiKey)
	}
	return url.Values(qs).Encode()
}

// symbolsToString converts an array of symbols to the format
// needed for a multiset Quandl request.
func symbolsToString(symbols []string) string {
	var result []string
	for _, symbol := range symbols {
		result = append(result, strings.Replace(symbol, "/", ".", -1))
	}
	return strings.Join(result, ",")
}

// marshallerError returns a formatted error that includes the response
func marshallerError(response []byte, err error) error {
	return errors.New("JSON Marshaller Error:\nRESPONSE:\n" + string(response) + "\n\nERROR:\n" + err.Error())
}
