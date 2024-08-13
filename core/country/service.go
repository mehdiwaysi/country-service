package country

import "context"

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
