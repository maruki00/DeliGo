package shared_valueobject

import "fmt"

type AddressValueObject struct {
	country string
	city    string
	street  string
	house   int
	flat    int
}

func (obj *AddressValueObject) New(
	country string,
	city string,
	street string,
	house int,
	flat int) {
	obj.country = country
	obj.city = city
	obj.street = street
	obj.house = house
	obj.flat = flat
}

func (obj *AddressValueObject) ToString() string {
	return fmt.Sprintf("%s, %s, %s, %d, %d.", obj.country, obj.city, obj.street, obj.house, obj.flat)
}
