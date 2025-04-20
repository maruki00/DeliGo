package usecases

import (
	"deligo/internal/iam/domain/contracts"
)

type GroupUseCase struct {
	groupRepo contracts.IGroupRepository
}

func NewGroupUseCase(groupRepo contracts.IGroupRepository) *GroupUseCase {
	return &GroupUseCase{
		groupRepo: groupRepo,
	}
}
