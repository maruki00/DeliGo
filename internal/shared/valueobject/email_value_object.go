package sharedvo

type EmailValueObject struct {
	value string
}

func (obj *EmailValueObject) SetEmail(email string) {
	obj.value = email
}

func (obj *EmailValueObject) ToString() string {
	return obj.value
}
