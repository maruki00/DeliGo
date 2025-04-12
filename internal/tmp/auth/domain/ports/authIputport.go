package ports

import (
	auth_domain_dtos "delivery/internal/auth/domain/dtos"
	shared_domain_contracts "delivery/internal/shared/domain/contracts"
)

type AuthInputPort interface {
	Login(dto auth_domain_dtos.LoginDTO) shared_domain_contracts.ViewModel
	Register(dto auth_domain_dtos.RegisterDTO) shared_domain_contracts.ViewModel
	TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) shared_domain_contracts.ViewModel
	Logout(dto auth_domain_dtos.LogoutDTO)
}
