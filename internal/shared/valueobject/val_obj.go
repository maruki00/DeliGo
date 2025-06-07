package sharedvo

type ValObj interface {
	SetValue(valu any)
	GetValue() any
	String() string
}
