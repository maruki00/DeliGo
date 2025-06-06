package shared_valueobject

import "fmt"

type AddressValueObject struct {
	country string
	city    string
	street  string
	house   int
	flat    int
}

func (obj *AddressValueObject) String() string {
	return fmt.Sprintf("%s, %s, %s, %d, %d.", obj.country, obj.city, obj.street, obj.house, obj.flat)
}
