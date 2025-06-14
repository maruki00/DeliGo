package sharedcontract

import sharedmodel "github.com/maruki00/deligo/internal/shared/model"

type ViewModel interface {
	GetResponse() sharedmodel.ResponseModel
	String() string
}
