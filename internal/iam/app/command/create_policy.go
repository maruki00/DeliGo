package command

import (
	"github.com/maruki00/deligo/internal/iam/infra/model"
)

type CreatePolicy struct {
	Role       string  // V0
	Resource   string  // V1
	Action     string  // V2
	Effect     *string // V3
	IpSource   *string // V4
	TimeWindow *string // V5
}

func (_this *CreatePolicy) Name() string {
	return "CreatePolicy"
}

func (_this *CreatePolicy) MapToModel() *model.Policy {
	return &model.Policy{
		Role:       _this.Role,
		Resource:   _this.Resource,
		Action:     _this.Action,
		Effect:     _this.Effect,
		IpSource:   _this.IpSource,
		TimeWindow: _this.TimeWindow,
	}
}
