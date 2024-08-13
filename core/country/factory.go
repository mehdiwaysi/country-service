package country

import "context"

type Factory interface {
	CreateCountryRepository(ctx context.Context) (Repository, error)
}
