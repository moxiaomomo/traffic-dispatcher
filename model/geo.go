package model

import "go.mongodb.org/mongo-driver/bson"

// GeoLocation 经纬度位置
type GeoLocation struct {
	// Lat 纬度
	Lat float64 `json:"lat"`
	// Lng 经度
	Lng float64 `json:"lng"`
}

// UserLocation : 用户位置结构体
type UserLocation struct {
	Name    string `json:"name" bson:"name"`
	UID     string `json:"uid" bson:"uid"`
	H3Index string `json:"h3index" bson:"h3index"`
	GeoInfo bson.M `json:"geoinfo" bson:"geoinfo"`
}
