Go Quandl
=========

This library provides easy access to the 
[Quandl API](https://www.quandl.com/help/api) 
using the Go programming language.

The full documentation is at:  
[godoc.org/github.com/DannyBen/quandl](http://godoc.org/github.com/DannyBen/quandl)

Install
-------

	$ go get github.com/DannyBen/quandl

Features
--------

* Supports 4 call types to Quandl: `GetSymbol`, `GetSymbols`, `GetList` and `GetSearch`
* Returns either a native Go object, or a raw (CSV/JSON/XML)
  response.
* Support for cache handling.

Usage
-----
Basic usage looks like this:

	quandl.ApiKey = "YOUR KEY"
	data, err := quandl.GetSymbol("WIKI/AAPL", nil)

and will return a native Go object. To use the data in the
response, iterate through its Data property:

	for i, item := range data.Data {
	    fmt.Println(i, item[0], item[2])
	}

To receive a raw response from Quandl (CSV, JSON, XML)
you can use:

	data, err := quandl.GetSymbolRaw("WIKI/AAPL", "csv", nil)

To pass options to the Quandl API, use something like this:

	v := quandl.Options{}
	v.Set("trim_start", "2014-01-01")
	v.Set("trim_end", "2014-02-02")
	data, err := quandl.GetSymbol("WIKI/AAPL", v)

More examples are in the [quandl_test file](https://github.com/DannyBen/quandl/blob/master/quandl_test.go)
