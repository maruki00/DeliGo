package shared_valueobject

type ValObj interface {
	SetValue(valu any)
	GetValue() any
	String() string
}
