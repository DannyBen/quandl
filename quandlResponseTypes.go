package quandl

// Type SymbolResponse represents the response from Quandl
// when requesting a single symbol
type SymbolResponse struct {
	Dataset
	Data [][]interface{}
}

// Type ListResponse represents the response received when
// requesting a list of supported symbol in a data source
type ListResponse struct {
	Datasets []Dataset
	Meta     ResponseMeta
}

// Type SearchResponse represents the response received
// when submitting a search request
type SearchResponse struct {
	ListResponse
}

// Type Document represents an entity at Quandl.
// It is used when requesting a symbol data, or a list of
// symbols supported by a data source.
type Dataset struct {
	Id                  uint
	DatasetCode         string `json:"dataset_code"`
	DatabaseCode        string `json:"database_code"`
	Name                string
	Description         string
	RefreshedAt         string   `json:"refreshed_at"`
	NewestAvailableDate string   `json:"newest_available_date"`
	OldestAvailableDate string   `json:"oldest_available_date"`
	ColumnNames         []string `json:"column_names"`
	Frequency           string
	Type                string
	Premium             bool
	Limit               int
	Transform           string
	ColumnIndex         int
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	Collapse            string
	Order               string
	DatabaseId          uint
}

type ResponseMeta struct {
	PerPage          int `json:"per_page"`
	Query            string
	CurrentPage      int  `json:"current_page"`
	PrevPage         int  `json:"prev_page"`
	TotalPages       uint `json:"total_pages"`
	TotalCount       uint `json:"total_count"`
	NextPage         int  `json:"next_page"`
	CurrentFirstItem int  `json:"current_first_item"`
	CurrentLastItem  int  `json:"current_last_item"`
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
