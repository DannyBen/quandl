// Package quandl provides easy access to the Quandl API
//
// It provides methods for getting response from Quandl in
// several formats.
//
// Basic usage looks like this:
//
//    quandl.ApiKey = "YOUR KEY"
//    data, err := quandl.GetSymbol("WIKI/AAPL", nil)
//
// and will return a native Go object. To use the data in the
// response, iterate through its Data property:
//
//    for i, item := range data.Data {
//        fmt.Println(i, item[0], item[2])
//    }
//
// To receive a raw response from Quandl (CSV, JSON, XML)
// you can use:
//
//    data, err := quandl.GetSymbolRaw("WIKI/AAPL", "csv", nil)
//
// To pass options to the Quandl API, use something like this:
//
//    v := quandl.Options{}
//    v.Set("trim_start", "2014-01-01")
//    v.Set("trim_end", "2014-02-02")
//    data, err := quandl.GetSymbol("WIKI/AAPL", v)
//
package quandl
