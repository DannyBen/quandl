package quandl_test

import (
	"fmt"
	"github.com/DannyBen/filecache"
	"github.com/DannyBen/quandl"
	"time"
)

var apiKey = "PUT_KEY_HERE"

// Main Functions

func ExampleGetSymbol() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")
	// ---

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
	// ---

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

func ExampleGetList() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := quandl.GetList("WIKI", 1, 3)
	if err != nil {
		panic(err)
	}

	for i, doc := range data.Docs {
		fmt.Println(i, doc.Code)
		break
	}

	// Output:
	// 0 AAPL
}

func ExampleGetSearch() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := quandl.GetSearch("google stock", 1, 3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %v results", len(data.Docs))

	// Output:
	// Found 3 results
}

// ToColumns Functions

func ExampleToColumns() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-06")
	v.Set("trim_end", "2014-01-08")
	v.Set("column", "4")
	// ---

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := quandl.ToColumns(data.Data)
	fmt.Println(d)

	// Output:
	// [[2014-01-08 2014-01-07 2014-01-06] [543.46 540.04 543.93]]
}

func ExampleToNamedColumns() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-06")
	v.Set("trim_end", "2014-01-07")
	v.Set("column", "4")
	// ---

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := quandl.ToNamedColumns(data.Data, data.ColumnNames)
	fmt.Println(d["Date"], d["Close"])

	// Output:
	// [2014-01-07 2014-01-06] [540.04 543.93]
}

// Column Converters

func ExampleFloatColumn() {
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}

	opts := quandl.NewOptions(
		"trim_start", "2014-01-01",
		"trim_end", "2014-01-06",
		"column", "4",
	)

	data, err := quandl.GetSymbol("WIKI/AAPL", opts)
	if err != nil {
		panic(err)
	}

	columns := data.ToColumns()
	var prices []float64 = quandl.FloatColumn(columns[1])
	fmt.Println(prices)

	// Output:
	// [543.93 540.98 553.13]
}

func ExampleTimeColumn() {
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}

	opts := quandl.NewOptions(
		"trim_start", "2014-01-01",
		"trim_end", "2014-01-06",
		"column", "4",
	)

	data, err := quandl.GetSymbol("WIKI/AAPL", opts)
	if err != nil {
		panic(err)
	}

	columns := data.ToColumns()
	var dates []time.Time = quandl.TimeColumn(columns[0])
	fmt.Println(dates)

	// Output:
	// [2014-01-06 00:00:00 +0000 UTC 2014-01-03 00:00:00 +0000 UTC 2014-01-02 00:00:00 +0000 UTC]
}

// Response Types

func ExampleSymbolResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")
	// ---

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
	// Wiki EOD Stock Prices
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

func ExampleListResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	// ---

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
	// STSI
}

func ExampleSearchResponse() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := quandl.GetSearch("twitter", 2, 5)
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
	fmt.Println(source.Description[:20], "...")
	fmt.Println(source.Name)
	fmt.Println(source.Host)
	fmt.Println(source.Premium)

	// Output:
	// Found more than 1000 results
	// 2
	// 5
	// [date Followers Following Favorites Tweets Listed]
	// <nil>
	// 13811288
	// Twitter Inc.
	// TWITTER
	// TO_BE
	// to be Twit ...
	// to-be-Twitter-Metrics
	// http://twi ...
	// Collage the internet ...
	// 201 ...
	// daily
	// 201 ...
	// 201 ...
	// false

	// false
	// 12832
	// TWITTER
	// 98506
	// Official Twitter sta ...
	// Twitter Inc.
	// twitter.com
	// false

}

// Response Types Functions

func ExampleSymbolResponse_ToColumns() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-06")
	v.Set("trim_end", "2014-01-08")
	v.Set("column", "4")
	// ---

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToColumns()
	fmt.Println(d)

	// Output:
	// [[2014-01-08 2014-01-07 2014-01-06] [543.46 540.04 543.93]]
}

func ExampleSymbolResponse_ToNamedColumns_1() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-06")
	v.Set("trim_end", "2014-01-07")
	v.Set("column", "11")
	// ---

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToNamedColumns(nil)
	fmt.Println(d["Date"], d["Adj. Close"])

	// Output:
	// [2014-01-07 2014-01-06] [75.561212341336 76.105492609478]
}

func ExampleSymbolResponse_ToNamedColumns_2() {
	// This block is optional
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}
	v := quandl.Options{}
	v.Set("trim_start", "2014-01-06")
	v.Set("trim_end", "2014-01-07")
	v.Set("column", "11")
	// ---

	data, err := quandl.GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToNamedColumns([]string{"date", "close"})
	fmt.Println(d["date"], d["close"])

	// Output:
	// [2014-01-07 2014-01-06] [75.561212341336 76.105492609478]
}

// Options

func ExampleNewOptions() {
	quandl.ApiKey = apiKey
	quandl.CacheHandler = filecache.Handler{Life: 60}

	opts := quandl.NewOptions(
		"trim_start", "2014-01-01",
		"trim_end", "2014-01-06",
		"column", "4",
	)

	data, err := quandl.GetSymbolRaw("WIKI/AAPL", "csv", opts)
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

// Naming convention cheatsheet
// Example()     Example_1()
// ExampleF()    ExampleF_1()
// ExampleT()    ExampleT_1()
// ExampleT_M()  ExampleT_M_1()
