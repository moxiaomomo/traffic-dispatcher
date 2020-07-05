package model

import "go.mongodb.org/mongo-driver/bson"

// Driver : 司机结构体
type Driver struct {
	Name    string    `json:"name"`
	UID     string    `json:"uid"`
	H3Index string    `json:"h3index"`
	GeoType string    `json:"type"`
	Coord   []float64 `json:"coordinates"`
	GeoInfo bson.M    `json:"geo"`
}
