package quandl

import (
	"github.com/DannyBen/filecache"

	"fmt"
	"os"
	"time"
)

func init() {
	ApiKey = os.Getenv("QUANDL_KEY")
}

// Main Functions
func ExampleGetSymbol() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-01")
	v.Set("end_date", "2014-02-02")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Symbol: %v, Row Count: %v\n", data.DatasetCode, len(data.Data))
	fmt.Printf("Fifth column is named %v\n", data.ColumnNames[4])
	fmt.Printf("On %v the close price was %v\n", data.Data[1][0], data.Data[1][4])

	// Output:
	// Symbol: AAPL, Row Count: 21
	// Fifth column is named Close
	// On 2014-01-30 the close price was 499.782
}

func ExampleGetSymbolRaw() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-01")
	v.Set("end_date", "2014-01-06")
	v.Set("column_index", "4") // Close price only
	// ---

	data, err := GetSymbolRaw("WIKI/AAPL", "csv", v)
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
	CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := GetList("WIKI", 1, 3)
	if err != nil {
		panic(err)
	}

	for i, doc := range data.Datasets {
		fmt.Println(i, doc.DatabaseCode, doc.ColumnNames[1])
	}

	// Output:
	// 0 WIKI Open
	// 1 WIKI Open
	// 2 WIKI Open
}

func ExampleGetSearch() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := GetSearch("google stock", 1, 3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %v results", len(data.Datasets))

	// Output:
	// Found 3 results
}

// ToColumns Functions

func ExampleToColumns() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-06")
	v.Set("end_date", "2014-01-08")
	v.Set("column_index", "4")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := ToColumns(data.Data)
	fmt.Println(d)

	// Output:
	// [[2014-01-08 2014-01-07 2014-01-06] [543.46 540.0375 543.93]]
}

func ExampleToNamedColumns() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-06")
	v.Set("end_date", "2014-01-07")
	v.Set("column_index", "4")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := ToNamedColumns(data.Data, data.ColumnNames)
	fmt.Println(d["Date"], d["Close"])

	// Output:
	// [2014-01-07 2014-01-06] [540.0375 543.93]
}

// Column Converters

func ExampleFloatColumn() {
	CacheHandler = filecache.Handler{Life: 60}

	opts := NewOptions(
		"start_date", "2014-01-01",
		"end_date", "2014-01-06",
		"column_index", "4",
	)

	data, err := GetSymbol("WIKI/AAPL", opts)
	if err != nil {
		panic(err)
	}

	columns := data.ToColumns()
	var prices []float64 = FloatColumn(columns[1])
	fmt.Println(prices)

	// Output:
	// [543.93 540.98 553.13]
}

func ExampleTimeColumn() {
	CacheHandler = filecache.Handler{Life: 60}

	opts := NewOptions(
		"start_date", "2014-01-01",
		"end_date", "2014-01-06",
		"column_index", "4",
	)

	data, err := GetSymbol("WIKI/AAPL", opts)
	if err != nil {
		panic(err)
	}

	columns := data.ToColumns()
	var dates []time.Time = TimeColumn(columns[0])
	fmt.Println(dates)

	// Output:
	// [2014-01-06 00:00:00 +0000 UTC 2014-01-03 00:00:00 +0000 UTC 2014-01-02 00:00:00 +0000 UTC]
}

// Response Types

func ExampleSymbolResponse() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-01")
	v.Set("end_date", "2014-02-02")
	// ---

	data, err := GetSymbol("WIKI/MSFT", v)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Id)
	fmt.Println(data.DatasetCode)
	fmt.Println(data.DatabaseCode)
	fmt.Println(data.Name)
	fmt.Println(data.Description[:20], "...")
	fmt.Println(data.RefreshedAt)
	fmt.Println(data.NewestAvailableDate)
	fmt.Println(data.OldestAvailableDate)
	fmt.Println(data.ColumnNames[0], "...")
	fmt.Println(data.Frequency)
	fmt.Println(data.Type)
	fmt.Println(data.Premium)
	fmt.Println(data.Limit)
	fmt.Println(data.Transform)
	fmt.Println(data.ColumnIndex)
	fmt.Println(data.StartDate)
	fmt.Println(data.EndDate)
	fmt.Println(data.Collapse)
	fmt.Println(data.Order)
	fmt.Println(data.DatabaseId)

	// TODO
	// Output:
	// Date ...
	// map[]
	// 9775827
	// Wiki EOD Stock Prices
	// WIKI
	// MSFT
	// Microsoft Corporatio ...
	// Microsoft-Corporatio ...
	// http://www.com/WIKI/MSFT
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
	CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := GetList("WIKI", 2, 5)
	if err != nil {
		panic(err)
	}

	if data.Meta.TotalCount > 3000 {
		fmt.Println("Found over 3000 results")
	}
	fmt.Println(data.Meta.CurrentPage)
	fmt.Println(data.Meta.PerPage)
	// fmt.Println(data.Docs[0].Code)

	// Output:
	// Found over 3000 results
	// 2
	// 5
}

func ExampleSearchResponse() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	// ---

	data, err := GetSearch("twitter", 2, 5)
	if err != nil {
		panic(err)
	}

	if data.Meta.TotalCount > 1000 {
		fmt.Println("Found more than 1000 results")
	}
	fmt.Println(data.Meta.CurrentPage)
	fmt.Println(data.Meta.PerPage)

	doc := data.Datasets[0]
	fmt.Println(doc.Id)
	fmt.Println(doc.DatasetCode)
	fmt.Println(doc.DatabaseCode)
	fmt.Println(doc.Name)
	fmt.Println(doc.Description[:20], "...")
	fmt.Println(doc.RefreshedAt)
	fmt.Println(doc.NewestAvailableDate)
	fmt.Println(doc.OldestAvailableDate)
	fmt.Println(doc.ColumnNames[0], "...")
	fmt.Println(doc.Frequency)
	fmt.Println(doc.Type)
	fmt.Println(doc.Premium)
	fmt.Println(doc.Limit)
	fmt.Println(doc.Transform)
	fmt.Println(doc.ColumnIndex)
	fmt.Println(doc.StartDate)
	fmt.Println(doc.EndDate)
	fmt.Println(doc.Collapse)
	fmt.Println(doc.Order)
	fmt.Println(doc.DatabaseId)

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
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-06")
	v.Set("end_date", "2014-01-08")
	v.Set("column_index", "4")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToColumns()
	fmt.Println(d)

	// Output:
	// [[2014-01-08 2014-01-07 2014-01-06] [543.46 540.0375 543.93]]
}

func ExampleSymbolResponse_ToNamedColumns_1() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-06")
	v.Set("end_date", "2014-01-07")
	v.Set("column_index", "11")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToNamedColumns(nil)
	fmt.Println(d["Date"], d["Adj. Close"])

	// Output:
	// [2014-01-07 2014-01-06] [73.451508457897 73.98093464899]
}

func ExampleSymbolResponse_ToNamedColumns_2() {
	// This block is optional
	CacheHandler = filecache.Handler{Life: 60}
	v := Options{}
	v.Set("start_date", "2014-01-06")
	v.Set("end_date", "2014-01-07")
	v.Set("column_index", "11")
	// ---

	data, err := GetSymbol("WIKI/AAPL", v)
	if err != nil {
		panic(err)
	}

	d := data.ToNamedColumns([]string{"date", "close"})
	fmt.Println(d["date"], d["close"])

	// Output:
	// [2014-01-07 2014-01-06] [73.451508457897 73.98093464899]
}

// Options

func ExampleNewOptions() {
	CacheHandler = filecache.Handler{Life: 60}

	opts := NewOptions(
		"start_date", "2014-01-01",
		"end_date", "2014-01-06",
		"column_index", "4",
	)

	data, err := GetSymbolRaw("WIKI/AAPL", "csv", opts)
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
