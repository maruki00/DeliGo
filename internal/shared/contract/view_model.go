package sharedcontract

import sharedmodel "deligo/internal/shared/model"

type ViewModel interface {
	GetResponse() sharedmodel.ResponseModel
	String() string
}
