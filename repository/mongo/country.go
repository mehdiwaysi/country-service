package mongo

import (
	"github.com/revotech-group/go-mongo/repository"
)

type CountryRepository struct {
	*repository.Repository[any]
}
