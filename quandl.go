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
	"time"
)

// ApiKey is used to set your api key before you make any call
var ApiKey = ""

// LastUrl will hold the last requested URL after each call
var LastUrl = ""

// CacheHandler is a reference to a struct that implements the Cacher interface.
// If set, it will use it to get documents from the cache or set to it.
var CacheHandler Cacher

var urlTemplates map[string]string = map[string]string{
	"symbol":  "https://www.quandl.com/api/v1/datasets/%s.%s?%s",
	"symbols": "https://www.quandl.com/api/v1/multisets.%s?columns=%s&%s",
	"search":  "https://www.quandl.com/api/v1/datasets.%s?%s",
	"list":    "http://www.quandl.com/api/v2/datasets.%s?%s",
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

// NewOptions accepts any even number of arguments and returns an
// Options object. The odd arguments are the keys, the even arguments
// are the values.
func NewOptions(s ...string) Options {
	o := Options{}
	for i := 0; i < len(s); i += 2 {
		o.Set(s[i], s[i+1])
	}
	return o
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

// ToColumns converts a rows array to a columns array
func ToColumns(src [][]interface{}) (out [][]interface{}) {
	out = make([][]interface{}, len(src[0]))
	for _, row := range src {
		for j, cell := range row {
			out[j] = append(out[j], cell)
		}
	}
	return
}

// ToNamedColumns converts a rows array to a columns map
func ToNamedColumns(src [][]interface{}, keys []string) (out map[string][]interface{}) {
	out = make(map[string][]interface{})
	for _, row := range src {
		for j, cell := range row {
			out[keys[j]] = append(out[keys[j]], cell)
		}
	}
	return
}

// FloatColumn converts a column of interface{} to a column of floats
func FloatColumn(s []interface{}) []float64 {
	r := make([]float64, len(s))
	for i := range s {
		r[i] = s[i].(float64)
	}
	return r
}

// TimeColumn converts a column of interface{} to a column of time
func TimeColumn(s []interface{}) []time.Time {
	r := make([]time.Time, len(s))
	for i := range s {
		r[i], _ = time.Parse("2006-01-02", s[i].(string))
	}
	return r
}

// StringColumn converts a column of interface{} to a column of string
func StringColumn(s []interface{}) []string {
	r := make([]string, len(s))
	for i := range s {
		r[i] = s[i].(string)
	}
	return r
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
