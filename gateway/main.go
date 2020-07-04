package main

import (
	"fmt"
	"net/http"
	"strconv"

	h3 "github.com/uber/h3-go/v3"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	lat, err1 := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, err2 := strconv.ParseFloat(r.Form.Get("lon"), 64)
	if err1 != nil || err2 != nil {
		w.Write([]byte("FAILED"))
		return
	}

	geo := h3.GeoCoord{
		Latitude:  lat, // 37.775938728915946,
		Longitude: lon, // -122.41795063018799,
	}
	resolution := 9

	result := fmt.Sprintf("%#x", h3.FromGeo(geo, resolution))
	// Output:
	// 0x8928308280fffff

	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
