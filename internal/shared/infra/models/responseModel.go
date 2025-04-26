package shared_models

import "encoding/json"

type ResponseModel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
	Result  any    `json:"result"`
}

func (obj *ResponseModel) ToMap() map[string]any {
	return map[string]any{
		"status":  obj.Status,
		"message": obj.Message,
		"error":   obj.Error,
		"result":  obj.Result,
	}
}

func (obj *ResponseModel) String() string {
	d, err := json.Marshal(obj)
	if err != nil {
		panic("could not marshal the responsemodel object")
	}
	return string(d)

}

func (obj *ResponseModel) ToString() string {
	d, err := json.Marshal(obj)
	if err != nil {
		panic("could not marshal the responsemodel object")
	}
	return string(d)

}
