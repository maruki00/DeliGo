package shared_valueobject

type ValueObject interface {
	SetValue(valu any)
	GetValue() any
	String() string
}
