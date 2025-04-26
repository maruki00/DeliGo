package userHandlers

import (
	"context"
	userCommands "deligo/internal/iam/app/user/commands"
	"deligo/internal/iam/domain/contracts"
	valueobjects "deligo/internal/iam/domain/valueobject"
	"deligo/internal/iam/infra/models"
	pkgCqrs "deligo/pkg/cqrs"
)

type CreateUserHandler struct {
	userRepo contracts.IUserRepository
}

func NewCreateUserHandler(userRepo contracts.IUserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		userRepo: userRepo,
	}
}

func (_this *CreateUserHandler) Handle(ctx context.Context, command pkgCqrs.Command) error {
	cmd := command.(*userCommands.CreateUserCommand)
	pwd, err := valueobjects.NewPassword(cmd.Password)
	if err != nil {
		return err
	}
	u := models.User{
		ID:                valueobjects.ID(cmd.ID.String()), //valueobjects.ID       `json:"id" gorm:"type:uuid;primaryKey"`
		Username:          cmd.Username,                     //string                `json:"user_name" gorm:"type:varchar(255)"`
		Email:             cmd.Email,                        //string                `json:"email" gorm:"type:varchar(255)"`
		TenantID:          valueobjects.ID("1234"),          //valueobjects.ID       `json:"tenant_id" gorm:"type:varchar(255)"`
		Password:          pwd,                              //valueobjects.Password `json:"password" gorm:"type:varchar(255)"`
		PasswordChangedAt: nil,                              //*time.Time            `json:"password_changed_at" gorm:"type:varchar(255)"`
		IsActive:          false,                            //bool                  `json:"is_active" gorm:"default:0"`
		LastLogin:         nil,                              //*time.Time            `json:"last_login" `
		MFAEnabled:        false,                            //bool                  `json:"mfa_enabled" gorm:"default:0"`
		MFASecret:         "",                               //string                `json:"mfa_secret"`
		Profile:           nil,                              //*Profile              `json:"profile" gorm:"foreignKey:UserID"`
		Groups:            nil,                              //[]*Group              `json:"groups" gorm:"many2many:user_groups;"`
		Policies:          nil,                              //[]*Policy             `json:"policies" gorm:"many2many:user_policies;"`
	}
	err = _this.userRepo.Save(ctx, &u)
	// 	&models.User{
	// 	ID:                valueobjects.ID(cmd.ID.String()),
	// 	Username:          cmd.Username,
	// 	Email:             cmd.Email,
	// 	Password:          valueobjects.Password(cmd.Password),
	// 	PasswordChangedAt: cmd.PasswordChangedAt,
	// 	IsActive:          cmd.IsActive,
	// 	MFAEnabled:        cmd.MFAEnabled,
	// 	MFASecret:         cmd.MFASecret,
	// })
	return err
}
