package model

import "go.mongodb.org/mongo-driver/bson"

// Driver : 司机结构体
type Driver struct {
	Name    string `bson:"name"`
	UID     string `bson:"uid"`
	H3Index string `bson:"h3index"`
	// GeoType string    `bson:"type"`
	// Coord   []float64 `bson:"coordinates"`
	GeoInfo bson.M `bson:"geoinfo"`
}
