package mongo

import "go.mongodb.org/mongo-driver/bson"

func generateSortQuery(sort string) bson.D {
	sortQuery := bson.D{}
	if sort != "" {
		startWithDash := string(sort[0]) == "-"
		if startWithDash {
			sortQuery = append(sortQuery, bson.E{Key: string(sort[1:]), Value: -1})
		} else {
			sortQuery = append(sortQuery, bson.E{Key: sort, Value: 1})
		}
	}
	return sortQuery
}
