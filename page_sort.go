package relpaginator

import (
	"encoding/json"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
)

type PageSort struct {
	Page         uint     `json:"page_number,omitempty"`
	ItemsPerPage uint     `json:"itemsPerPage,omitempty"`
	SortBy       []string `json:"sortBy,omitempty"`
	SortDesc     []bool   `json:"sortDesc,omitempty"`
	GroupBy      []string `json:"groupBy,omitempty"`
	GroupDesc    []bool   `json:"groupDesc,omitempty"`
	// TODO: make mustSort and multiSort useful
	MustSort  bool          `json:"mustSort,omitempty"`
	MultiSort bool          `json:"multiSort,omitempty"`
	filters   []rel.Querier `json:"-"`
}

func CreatePageSortFromMap(values map[string]interface{}) *PageSort {
	var res PageSort

	b, e := json.Marshal(values)
	if e != nil {
		panic(e)
	}
	if e := json.Unmarshal(b, &res); e != nil {
		panic(e)
	}
	return &res
}

// SetFiltersQueries allows to add filters to be consider to the pagination
func (c *PageSort) SetFiltersQueries(queries ...rel.Querier) {
	c.filters = queries
}

func (c *PageSort) GetFiltersQueries() []rel.Querier {
	return c.filters
}

func (c *PageSort) GetSortQueries() []rel.Querier {
	var sortQueries []rel.Querier
	for i := 0; i < len(c.SortBy); i++ {
		var sortQuery rel.Querier
		isSortDesc := c.SortDesc[i]
		if isSortDesc {
			sortQuery = sort.Desc(c.SortBy[i])
		} else {
			sortQuery = sort.Asc(c.SortBy[i])
		}

		sortQueries = append(sortQueries, sortQuery)
	}
	return sortQueries
}

func (c *PageSort) GetItemsPerPage() uint {
	return uint(c.ItemsPerPage)
}
func (c *PageSort) GetPage() uint {
	return uint(c.Page)
}
