package shared_valueobject

type FloafVAlueObject32 struct {
	value float32
}

func (obj *FloafVAlueObject32) NewFloafVAlueObject(value float32) {
	obj.value = value
}

func (obj FloafVAlueObject32) getValue() float32 {
	return obj.value
}
