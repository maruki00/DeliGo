package shared_valueobject

import (
	"fmt"
)

type Cordinate struct {
	lat  float32
	long float32
}

func NewCordinate(lat, long float32) *Cordinate {
	if lat < 0.0 || long < 0.0 {
		panic("invalid lat or lang ")
	}

	return &Cordinate{
		lat:  lat,
		long: long,
	}
}

func (obj Cordinate) String() string {
	return fmt.Sprintf("%f,%f", obj.lat, obj.long)
}
