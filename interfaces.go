package relpaginator

import (
	"context"
)

type Paginable interface {
	CreatePagination(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageSort *PageSort,
	) (*Paginator, error)
}
