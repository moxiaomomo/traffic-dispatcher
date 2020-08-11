package model

// GeoLocation 经纬度位置
type GeoLocation struct {
	// Lat 纬度
	Lat float64 `json:"lat"`
	// Lng 经度
	Lng float64 `json:"lng"`
}
