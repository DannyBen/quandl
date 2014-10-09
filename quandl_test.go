package quandl_test

import (
	"fmt"
	"github.com/DannyBen/quandl"
	"github.com/DannyBen/quandl/cache"
)

var apiKey = "PUT_KEY_HERE"

func ExampleGetSymbol() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = cache.Handler{}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Symbol: %v, Row Count: %v\n", data.Code, len(data.Data))
	fmt.Printf("Fifth column is named %v\n", data.ColumnNames[4])
	fmt.Printf("On %v the close price was %v\n", data.Data[1][0], data.Data[1][4])
	// Output:
	// Symbol: AAPL, Row Count: 21
	// Fifth column is named Close
	// On 2014-01-30 the close price was 499.78
}

func ExampleGetSymbolRaw() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = cache.Handler{}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-01-06")
	v.Set("column", "4") // Close price only

	data, err := quandl.GetSymbolRaw("WIKI/AAPL", "csv", v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// Output:
	// Date,Close
	// 2014-01-06,543.93
	// 2014-01-03,540.98
	// 2014-01-02,553.13
}

func ExampleGetSymbols() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = cache.Handler{}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-01-06")
	v.Set("sort_order", "asc")

	// Get two symbols at once, only close prices (column 4)
	symbols := []string{"WIKI/AAPL.4", "WIKI/CSCO.4"}

	data, err := quandl.GetSymbols(symbols, v)
	if err != nil {
		panic(err)
	}

	for i, row := range data.Data {
		fmt.Printf("Row:%v Date:%v AAPL:%v CSCO:%v\n", i, row[0], row[1], row[2])
	}

	// Output:
	// Row:0 Date:2014-01-02 AAPL:553.13 CSCO:22
	// Row:1 Date:2014-01-03 AAPL:540.98 CSCO:21.98
	// Row:2 Date:2014-01-06 AAPL:543.93 CSCO:22.01
}

func ExampleGetList() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = cache.Handler{}

	data, err := quandl.GetList("WIKI", 1, 3)
	if err != nil {
		panic(err)
	}

	for i, doc := range data.Docs {
		fmt.Println(i, doc.Code)
	}

	// Output:
	// 0 AAPL
	// 1 ATMI
	// 2 PACR
}

func ExampleGetSearch() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = cache.Handler{}

	data, err := quandl.GetSearch("google stock", 1, 3)
	if err != nil {
		panic(err)
	}

	for i, doc := range data.Docs {
		fmt.Println(i, doc.SourceName, doc.Code)
	}

	// Output:
	// 0 Damodaran Financial Data GOOG_ROE
	// 1 Chris-Freeman GOOGL
	// 2 CBOE Futures Exchange INDEX_VXGOG
}
