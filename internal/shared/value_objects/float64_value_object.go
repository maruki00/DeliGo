package shared_valueobject

type FloafVAlueObject64 struct {
	value float64
}

func (obj *FloafVAlueObject64) NewFloafVAlueObject(value float64) {
	obj.value = value
}

func (obj FloafVAlueObject64) getValue() float64 {
	return obj.value
}
