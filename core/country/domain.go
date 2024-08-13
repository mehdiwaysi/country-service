package country

import "errors"

// #TODO: add all other bson tags
// #TODO: add all other fields
type Country struct {
	Alpha2Code   string   `bson:"alpha2Code"`
	AltSpellings []string `bson:"altSpellings"`
	Area         float64
	Population   int
}

// validate function
func (c *Country) Validate() error {
	if c.Alpha2Code == "" {
		return errors.New(ErrAlpha2CodeRequired)
	}

	return nil
}
