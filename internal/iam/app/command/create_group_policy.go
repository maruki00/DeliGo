package command

import "github.com/maruki00/deligo/internal/iam/infra/model"

type CreateGroupPolicy struct {
	User   string
	Group  string
	Domain string
}

func (_this *CreateGroupPolicy) Name() string {
	return "CreateGroupPolicy"
}

func (_this *CreateGroupPolicy) MapToModel() *model.GroupPolicy {
	return &model.GroupPolicy{
		User:   _this.User,
		Group:  _this.Group,
		Domain: _this.Domain,
	}
}
