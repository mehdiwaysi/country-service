package mongo

import (
	"context"
	"fmt"
	"regexp"

	domain "github.com/mehdiwaysi/country-service/core/country"
	"github.com/revotech-group/go-lib/db"
	gomongo "github.com/revotech-group/go-mongo/core"
	"github.com/revotech-group/go-mongo/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *CountryRepository) ListLightCountries(ctx context.Context, query domain.Query) (*domain.LightQueryResult, error) {
	q := bson.M{}

	if query.Search != "" {
		q["$or"] = []bson.M{
			{"alpha3Code": bson.M{"$regex": regexp.QuoteMeta(query.Search), "$options": "i"}},
			{"name": bson.M{"$regex": regexp.QuoteMeta(query.Search), "$options": "i"}},
		}
	}

	opts := options.Find().SetProjection(bson.M{
		"alpha2Code": 1,
		"alpha3Code": 1,
		"name":       1,
	})

	res, err := r.Repository.FindWithPagination(ctx, q, query.Page, query.Limit, generateSortQuery(query.Order), opts)
	if err != nil {
		return nil, err
	}

	return &domain.LightQueryResult{
		PaginationResult: res,
	}, nil
}


func (r *CountryRepository) GetByCountryAlphaCode(ctx context.Context, alphaCode string) (*domain.Country, error) {
	q := bson.M{
		"$or": []bson.M{
			{"alpha2Code": alphaCode},
			{"alpha3Code": alphaCode},
		},
	}

	res, err := r.Repository.FindOne(ctx, q)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("Country does not exist")
	}

	return res, nil
}

func (r *CountryRepository) GetByCountryCapital(ctx context.Context, capital string) (*domain.Country, error) {
	q := bson.M{"capital": capital}

	res, err := r.Repository.FindOne(ctx, q)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("Country does not exist")
	}

	return res, nil
}

func (r *CountryRepository) GetByCountryName(ctx context.Context, name string) (*domain.Country, error) {
	q := bson.M{"name": name}

	res, err := r.Repository.FindOne(ctx, q)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("Country does not exist")
	}

	return res, nil
}

func (r *CountryRepository) GetByCountryCurrency(ctx context.Context, currency string) ([]domain.Country, error) {
	q := bson.M{"currencies.code": currency}

	res, err := r.Repository.Find(ctx, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CountryRepository) GetByCountryLang(ctx context.Context, lang string) ([]domain.Country, error) {
	q := bson.M{"languages.iso639_1": lang}

	res, err := r.Repository.Find(ctx, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CountryRepository) GetByCountryRegion(ctx context.Context, region string) ([]domain.Country, error) {
	q := bson.M{"region": region}

	res, err := r.Repository.Find(ctx, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// func (r *CountryRepository) CreateCountry(ctx context.Context, country domain.Country) (*domain.Country, error) {
// 	res, err := r.Repository.InsertOne(ctx, country)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }
