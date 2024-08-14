package country

import "time"

type Currency struct {
	Code   string `bson:"code"`
	Name   string `bson:"name"`
	Symbol string `bson:"symbol"`
}

type Language struct {
	Iso639_1   string `bson:"iso639_1"`
	Iso639_2   string `bson:"iso639_2"`
	Name       string `bson:"name"`
	NativeName string `bson:"nativeName"`
}

type LightCountry struct {
	Alpha2Code string `bson:"alpha2Code"`
	Alpha3Code string `bson:"alpha3Code"`
	Name       string `bson:"name"`
}

type RegionalBloc struct {
	Acronym       string   `bson:"acronym"`
	Name          string   `bson:"name"`
	OtherAcronyms []string `bson:"otherAcronyms"`
	OtherNames    []string `bson:"otherNames"`
}

type Translation struct {
	Br string `bson:"br"`
	De string `bson:"de"`
	Es string `bson:"es"`
	Fa string `bson:"fa"`
	Fr string `bson:"fr"`
	Hr string `bson:"hr"`
	It string `bson:"it"`
	Ja string `bson:"ja"`
	Nl string `bson:"nl"`
	Pt string `bson:"pt"`
}
type Country struct {
	ID             string         `bson:"id"`
	TenantAlias    string         `bson:"-"`
	Alpha2Code     string         `bson:"alpha2Code"`
	Alpha3Code     string         `bson:"alpha3Code"`
	AltSpellings   []string       `bson:"altSpellings"`
	Area           float64        `bson:"area"`
	Borders        []string       `bson:"borders"`
	CallingCodes   []string       `bson:"callingCodes"`
	Capital        string         `bson:"capital"`
	Cioc           string         `bson:"cioc"`
	Currencies     []Currency     `bson:"currencies"`
	Demonym        string         `bson:"demonym"`
	Flag           string         `bson:"flag"`
	Gini           float64        `bson:"gini"`
	Languages      []Language     `bson:"languages"`
	Latlng         []float64      `bson:"latlng"`
	Name           string         `bson:"name"`
	NativeName     string         `bson:"nativeName"`
	NumericCode    string         `bson:"numericCode"`
	Population     int            `bson:"population"`
	Region         string         `bson:"region"`
	RegionalBlocs  []RegionalBloc `bson:"regionalBlocs"`
	Subregion      string         `bson:"subregion"`
	Timezones      []string       `bson:"timezones"`
	TopLevelDomain []string       `bson:"topLevelDomain"`
	Translations   []Translation  `bson:"translations"`
	CreatedAt      *time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt      *time.Time     `bson:"updatedAt,omitempty"`
}
