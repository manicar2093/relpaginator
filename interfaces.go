package relpaginator

import (
	"context"
)

type Paginable interface {
	CreatePage(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageSort *PageSort,
	) (*Page, error)
}
