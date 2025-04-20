package usecases

import (
	"deligo/internal/iam/domain/contracts"
)

type PermissionUseCase struct {
	permissionRepo contracts.IPermissionRepository
}

func NewPermissionUseCase(permissionRepo contracts.IPermissionRepository) *PermissionUseCase {
	return &PermissionUseCase{
		permissionRepo: permissionRepo,
	}
}
