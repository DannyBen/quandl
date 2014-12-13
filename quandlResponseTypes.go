package quandl

// Type SymbolResponse represents the response from Quandl
// when requesting a single symbol
type SymbolResponse struct {
	Document
	Data [][]interface{}
}

// Type ListResponse represents the response received when
// requesting a list of supported symbol in a data source
type ListResponse struct {
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Docs        []Document
}

// Type SearchResponse represents the response received
// when submitting a search request
type SearchResponse struct {
	ListResponse
	Sources []Source
}

// Type Document represents an entity at Quandl.
// It is used when requesting a symbol data, or a list of
// symbols supported by a data source.
type Document struct {
	ColumnNames []string `json:"column_names"`
	Errors      interface{}
	Id          int
	SourceName  string `json:"source_name"`
	SourceCode  string `json:"source_code"`
	Code        string
	Name        string
	UrlizeName  string `json:"urlize_name"`
	DisplayUrl  string `json:"display_url"`
	Description string
	UpdatedAt   string `json:"updated_at"`
	Frequency   string
	FromDate    string `json:"from_date"`
	ToDate      string `json:"to_date"`
	Private     bool
	Type        string
	Premium     bool
}

// Type Source represents a data source.
// Used by Search Response
type Source struct {
	Id            int
	Code          string
	DataSetsCount int `json:"datasets_count"`
	Description   string
	Name          string
	Host          string
	Premium       bool
}

// ToColumns converts the data array to a columns array
func (s *SymbolResponse) ToColumns() [][]interface{} {
	return ToColumns(s.Data)
}

// ToNamedColumns converts the rows array to a columns map.
// You may provide it with an array of strings for the map keys
// (header columns), or nil to use the column names as returned
// from Quandl.
func (s *SymbolResponse) ToNamedColumns(keys []string) map[string][]interface{} {
	if keys == nil {
		keys = s.ColumnNames
	}
	return ToNamedColumns(s.Data, keys)
}
