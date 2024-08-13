package country

import (
	"context"

	"github.com/revotech-group/go-lib/db/record"
)

type Query struct {
	Page  int64
	Limit int64
	Order string
	Query any
}

type QueryResult struct {
	*record.PaginationResult[Country]
}

type Repository interface {
	ListCountries(ctx context.Context, query Query) (*QueryResult, error)
}
