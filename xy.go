package goods

import (
	"github.com/umahmood/haversine"
)

// in meter
func XYDistance(x1 float64, y1 float64, x2 float64, y2 float64) (dis float64) {
	co1 := haversine.Coord{Lat: y1, Lon: x1}
	co2 := haversine.Coord{Lat: y2, Lon: x2}
	_, km := haversine.Distance(co1, co2)
	return km * 1000
}

// in meter
func XYDistanceMeter(x1 float64, y1 float64, x2 float64, y2 float64) (dis int) {
	return RoundFloat64(XYDistance(x1, y1, x2, y2))
}
