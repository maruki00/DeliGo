package shared_valueobject

type StringValeObject struct {
	value string
}

func (obj *StringValeObject) NewStringValeObject(value string) {
	obj.value = value
}

func (obj StringValeObject) toString() string {
	return obj.value
}
