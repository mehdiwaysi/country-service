package mongo

import (
	"context"
	"regexp"

	domain "github.com/mehdiwaysi/country-service/core/country"
	"github.com/revotech-group/go-lib/db"
	gomongo "github.com/revotech-group/go-mongo/core"
	"github.com/revotech-group/go-mongo/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type CountryRepository struct {
	*repository.Repository[domain.Country]
}

func NewCountryRepository(ctx context.Context, connectionProvider db.ConnectionProvider) (*CountryRepository, error) {

	schemaOptions := gomongo.NewSchemaOptions().SetAutogenAlgorithm(gomongo.UUID_AlgorithmType).SetIdFieldName("ID")

	r, err := repository.NewRepository[domain.Country](
		"countries",
		connectionProvider,
		schemaOptions,
	)
	if err != nil {
		return nil, err
	}

	repo := &CountryRepository{
		Repository: r,
	}

	if err := repo.Connect(ctx); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *CountryRepository) ListCountries(ctx context.Context, query domain.Query) (*domain.QueryResult, error) {
	q := bson.M{}

	if query.Search != "" {
		q["$or"] = []bson.M{
			{"alpha3Code": bson.M{"$regex": regexp.QuoteMeta(query.Search), "$options": "i"}},
			{"name": bson.M{"$regex": regexp.QuoteMeta(query.Search), "$options": "i"}},
		}
	}

	res, err := r.Repository.FindWithPagination(ctx, q, query.Page, query.Limit, generateSortQuery(query.Order))
	if err != nil {
		return nil, err
	}

	return &domain.QueryResult{
		PaginationResult: res,
	}, nil
}
