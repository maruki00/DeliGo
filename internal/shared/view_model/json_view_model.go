package shared_viewmodel

import (
	shared_model "github.com/maruki00/deligo/internal/shared/infra/model"
	"encoding/json"
)

type JsonViewModel struct {
	Data shared_model.ResponseModel
}

func NewJsonViewModel(data shared_model.ResponseModel) JsonViewModel {
	return JsonViewModel{
		Data: data,
	}
}

func (obj JsonViewModel) GetResponse() shared_model.ResponseModel {
	return obj.Data
}

func (obj JsonViewModel) String() string {
	d, err := json.Marshal(obj.Data)
	if err != nil {
		return "error data could not be serialized"
	}
	return string(d)
}
