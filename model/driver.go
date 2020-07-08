package model

import "go.mongodb.org/mongo-driver/bson"

// Driver : 司机结构体
type Driver struct {
	Name    string `json:"name" bson:"name"`
	UID     string `json:"uid" bson:"uid"`
	H3Index string `json:"h3index" bson:"h3index"`
	GeoInfo bson.M `json:"geoinfo" bson:"geoinfo"`
}
