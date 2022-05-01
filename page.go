package relpaginator

type Paginator struct {
	TotalPages   uint        `json:"total_pages,omitempty"`
	CurrentPage  uint        `json:"current_page,omitempty"`
	PreviousPage uint        `json:"previous_page,omitempty"`
	PageSize     uint        `json:"page_size,omitempty"`
	NextPage     uint        `json:"next_page,omitempty"`
	TotalEntries uint        `json:"total_entries,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
