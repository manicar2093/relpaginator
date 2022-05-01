package relpaginator

import (
	"context"
	"reflect"

	"github.com/go-rel/rel"
)

type PaginableImpl struct {
	repo   rel.Repository
	config *RelPaginatorConfig
}

func NewPaginator(repo rel.Repository, config *RelPaginatorConfig) Paginable {
	return PaginableImpl{repo: repo, config: config}
}

func (c PaginableImpl) CreatePagination(
	ctx context.Context,
	tableName string,
	holder interface{},
	pageSort *PageSort,
) (*Paginator, error) {
	pageSize := c.getPageSize(pageSort)
	page := pageSort.GetPage()

	totalEntries, err := c.repo.Count(ctx, tableName, pageSort.GetFiltersQueries()...)
	if err != nil {
		return nil, err
	}

	totalPages := uint(totalEntries) / c.config.PageSize

	if totalEntries > 0 && uint(totalEntries) < c.config.PageSize {
		totalPages = 1
	}

	if page > totalPages {
		return nil, PageError{PageNumber: page}
	}

	pageLimitQuery := rel.Limit(pageSize)
	pageOffsetQuery := rel.Offset(createOffsetValue(pageSize, page))
	sorting := pageSort.GetSortQueries()
	queries := pageSort.GetFiltersQueries()

	queries = append(queries, sorting...)
	queries = append(queries, pageLimitQuery, pageOffsetQuery)

	err = c.repo.FindAll(ctx, holder, queries...)
	if err != nil {
		return nil, err
	}

	return &Paginator{
		TotalPages:   totalPages,
		CurrentPage:  page,
		PreviousPage: page - 1,
		NextPage:     calculateNextPage(page, totalPages),
		Data:         holder,
		TotalEntries: getSliceCount(holder),
		PageSize:     c.config.PageSize,
	}, nil
}

func createOffsetValue(pageSize, pageNumber uint) uint {
	return pageSize * (pageNumber - 1)
}

func calculateNextPage(pageNumber, totalPages uint) uint {
	if pageNumber == totalPages {
		return 1
	}
	return pageNumber + 1
}

func getSliceCount(data interface{}) uint {
	val := reflect.ValueOf(data)

	isSlice := func(val reflect.Value) uint {
		if val.Kind() == reflect.Slice {
			return uint(val.Len())
		}
		return 0
	}
	if isPtr := val.Kind() == reflect.Ptr; isPtr {
		return isSlice(reflect.Indirect(val))
	}

	return isSlice(val)

}

func (c PaginableImpl) getPageSize(pageSort *PageSort) uint {
	if pageSort.GetItemsPerPage() > 0 {
		return pageSort.GetItemsPerPage()
	}
	return c.config.PageSize
}
