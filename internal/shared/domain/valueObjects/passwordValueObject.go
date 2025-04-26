package shared_valueobject

type PasswordValueObject struct {
	value string
}

func (obj *PasswordValueObject) SetPassword(password string) {
	obj.value = password
}

func (obj *PasswordValueObject) ToString() string {
	return obj.value
}

func (obj *PasswordValueObject) toString() string {
	return obj.value
}
