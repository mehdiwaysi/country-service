package factory

import (
	"context"

	domain "github.com/mehdiwaysi/country-service/core/country"
	mongo "github.com/mehdiwaysi/country-service/repository/mongo"
)

type Factory interface {
	CreateCountryRepository(ctx context.Context) (domain.Repository, error)
}

type factory struct {
	tenantAlias string
}

func NewFactory(tenantAlias string) Factory {
	return &factory{
		tenantAlias: tenantAlias,
	}
}

func (f *factory) CreateCountryRepository(ctx context.Context) (domain.Repository, error) {
	return mongo.NewCountryRepository(ctx, NewConnectionProvider())
}
