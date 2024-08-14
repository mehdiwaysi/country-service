package country

import (
	"context"
)

type service struct {
	factory Factory
	repo    Repository
}

func NewService(ctx context.Context, factory Factory) (*service, error) {
	repo, err := factory.CreateCountryRepository(ctx)
	if err != nil {
		return nil, err
	}

	return &service{
		factory: factory,
		repo:    repo,
	}, nil
}

func (s *service) ListCountries(ctx context.Context, query Query) (*QueryResult, error) {
	return s.repo.ListCountries(ctx, query)
}

func (s *service) ListLightCountries(ctx context.Context, query Query) (*LightQueryResult, error) {
	return s.repo.ListLightCountries(ctx, query)
}

func (s *service) GetByCountryAlphaCode(ctx context.Context, alphaCode string) (*Country, error) {
	return s.repo.GetByCountryAlphaCode(ctx, alphaCode)
}

func (s *service) GetByCountryCapital(ctx context.Context, capital string) (*Country, error) {
	return s.repo.GetByCountryCapital(ctx, capital)
}

func (s *service) GetByCountryName(ctx context.Context, name string) (*Country, error) {
	return s.repo.GetByCountryName(ctx, name)
}

func (s *service) GetByCountryCurrency(ctx context.Context, currency string) ([]Country, error) {
	return s.repo.GetByCountryCurrency(ctx, currency)
}

func (s *service) GetByCountryLang(ctx context.Context, lang string) ([]Country, error) {
	return s.repo.GetByCountryLang(ctx, lang)
}

func (s *service) GetByCountryRegion(ctx context.Context, region string) ([]Country, error) {
	return s.repo.GetByCountryRegion(ctx, region)
}

func (s *service) CreateCountry(ctx context.Context, country Country) (*Country, error) {
	return s.repo.CreateCountry(ctx, country)
}
