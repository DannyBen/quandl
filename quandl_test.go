package quandl_test

import (
	"fmt"
	"github.com/DannyBen/filecache"
	"github.com/DannyBen/quandl"
)

var apiKey = "PUT_KEY_HERE"

func ExampleGetSymbol() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
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
	quandl.CacheHandler = filecache.Handler{Life: 60}
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
	quandl.CacheHandler = filecache.Handler{Life: 60}
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
	quandl.CacheHandler = filecache.Handler{Life: 60}

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
	quandl.CacheHandler = filecache.Handler{Life: 60}

	data, err := quandl.GetSearch("google stock", 1, 3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %v results", len(data.Docs))

	// Output:
	// Found 3 results
}

func ExampleSymbolResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")

	data, err := quandl.GetSymbol("WIKI/MSFT", v)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.ColumnNames[0], "...")
	fmt.Println(data.Errors)
	fmt.Println(data.Id)
	fmt.Println(data.SourceName)
	fmt.Println(data.SourceCode)
	fmt.Println(data.Code)
	fmt.Println(data.Name[:20], "...")
	fmt.Println(data.UrlizeName[:20], "...")
	fmt.Println(data.DisplayUrl)
	fmt.Println(data.Description[:20], "...")
	fmt.Println(data.UpdatedAt[:3], "...")
	fmt.Println(data.Frequency)
	fmt.Println(data.FromDate)
	fmt.Println(data.ToDate[:3], "...")
	fmt.Println(data.Private)
	fmt.Println(data.Type)
	fmt.Println(data.Premium)
	fmt.Println(data.Data[0][1])

	// Output:
	// Date ...
	// map[]
	// 9775827
	// Quandl Open Data
	// WIKI
	// MSFT
	// Microsoft Corporatio ...
	// Microsoft-Corporatio ...
	// http://www.quandl.com/WIKI/MSFT
	// End of day open, hig ...
	// 201 ...
	// daily
	// 1986-03-13
	// 201 ...
	// false
	//
	// false
	// 36.95
}

func ExampleSymbolsResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")

	data, err := quandl.GetSymbols([]string{"WIKI/MSFT", "WIKI/AAPL"}, v)
	if err != nil {
		panic(err)
	}

	fmt.Println(data.ColumnNames[1], "...")
	fmt.Println(data.Data[0][1])
	fmt.Println(data.Columns[2], "...")
	fmt.Println(data.Errors)
	fmt.Println(data.Frequency)
	fmt.Println(data.FromDate)
	fmt.Println(data.ToDate)

	// Output:
	// WIKI.MSFT - Open ...
	// 36.95
	// High ...
	// map[]
	// daily
	// 2014-01-02
	// 2014-01-31
}

func ExampleListResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}

	data, err := quandl.GetList("WIKI", 2, 5)
	if err != nil {
		panic(err)
	}

	if data.TotalCount > 3000 {
		fmt.Println("Found over 3000 results")
	}
	fmt.Println(data.CurrentPage)
	fmt.Println(data.PerPage)
	fmt.Println(data.Docs[0].Code)

	// Output:
	// Found over 3000 results
	// 2
	// 5
	// SCBT
}

func ExampleSearchResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}

	data, err := quandl.GetSearch("facebook", 2, 5)
	if err != nil {
		panic(err)
	}

	if data.TotalCount > 1000 {
		fmt.Println("Found more than 1000 results")
	}
	fmt.Println(data.CurrentPage)
	fmt.Println(data.PerPage)

	doc := data.Docs[0]
	fmt.Println(doc.ColumnNames)
	fmt.Println(doc.Errors)
	fmt.Println(doc.Id)
	fmt.Println(doc.SourceName)
	fmt.Println(doc.SourceCode)
	fmt.Println(doc.Code)
	fmt.Println(doc.Name[:10], "...")
	fmt.Println(doc.UrlizeName)
	fmt.Println(doc.DisplayUrl[:10], "...")
	fmt.Println(doc.Description[:20], "...")
	fmt.Println(doc.UpdatedAt[:3], "...")
	fmt.Println(doc.Frequency)
	fmt.Println(doc.FromDate[:3], "...")
	fmt.Println(doc.ToDate[:3], "...")
	fmt.Println(doc.Private)
	fmt.Println(doc.Type)
	fmt.Println(doc.Premium)

	source := data.Sources[0]
	fmt.Println(source.Id)
	fmt.Println(source.Code)
	fmt.Println(source.DataSetsCount)
	fmt.Println(source.Description)
	fmt.Println(source.Name)
	fmt.Println(source.Host)
	fmt.Println(source.Premium)

	// Output:
	// Found more than 1000 results
	// 2
	// 5
	// [Date Capital Expenditures]
	// <nil>
	// 4417566
	// Damodaran Financial Data
	// DMDRN
	// FB_CAPEX
	// Facebook I ...
	// Facebook-Inc-FB-Capital-Expenditures
	// http://pag ...
	// This is the cumulate ...
	// 201 ...
	// annual
	// 201 ...
	// 201 ...
	// false
	//
	// false
	// 6946
	// DMDRN
	// 878156
	//
	// Damodaran Financial Data
	// pages.stern.nyu.edu/~adamodar/
	// false
}
