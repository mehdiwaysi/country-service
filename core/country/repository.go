package country

import (
	"context"

	"github.com/revotech-group/go-lib/db/record"
)

type Query struct {
	Page   int64
	Limit  int64
	Order  string
	Search string
}

type QueryResult struct {
	*record.PaginationResult[Country]
}

type LightQueryResult struct {
	*record.PaginationResult[LightCountry]
}

type Repository interface {
	ListCountries(ctx context.Context, query Query) (*QueryResult, error)
	// ListLightCountries(ctx context.Context, query Query) (*LightQueryResult, error)

	// GetByCountryAlphaCode(ctx context.Context, alphaCode string) (*Country, error)
	// GetByCountryCapital(ctx context.Context, capital string) (*Country, error)
	// GetByCountryName(ctx context.Context, name string) (*Country, error)
	// GetByCountryCurrency(ctx context.Context, currency string) ([]Country, error)
	// GetByCountryLang(ctx context.Context, lang string) ([]Country, error)
	// GetByCountryRegion(ctx context.Context, region string) ([]Country, error)

	// CreateCountry(ctx context.Context, country Country) (*Country, error)
}
