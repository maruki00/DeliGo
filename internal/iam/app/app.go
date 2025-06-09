package app

import (
	"context"
	"deligo/cmd/iam/configs"
	userCommand "deligo/internal/iam/app/user/command"
	userHandler "deligo/internal/iam/app/user/handler"
	userQuery "deligo/internal/iam/app/user/query"
	"deligo/internal/iam/app/user/service"
	contract "deligo/internal/iam/domain/contract"
	"deligo/internal/iam/infra/repository"
	pkgCqrs "deligo/pkg/cqrs"
	pkgPostgres "deligo/pkg/postgres"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db             *pkgPostgres.PGHandler
	UserRepo       contract.IUserRepository
	PermissionRepo contract.IPermissionRepository
	PolicyRepo     contract.IPolicyRepository
	UserServerSvc  *service.UserServerService
	UserCommandBus *pkgCqrs.CommandBus
	userQuerydBus  *pkgCqrs.QueryBus
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *configs.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.Postgres.Dsn)
	if err != nil {
		return nil, func() {}, err
	}
	userCommandBus := pkgCqrs.NewCommandBus()
	userQuerydBus := pkgCqrs.NewQueryBus()

	userRepo := repository.NewUserRepository(db)

	//permissionRepo := repository.NewPermissionRepository()

	userServiceSvc := service.NewUserUseCase(userCommandBus, userQuerydBus)

	userCommandBus.Register(&userCommand.CreateUserCommand{}, userHandler.NewCreateUserHandler(userRepo))
	userCommandBus.Register(&userCommand.DeleteUserCommand{}, userHandler.NewDeleteUserHandler(userRepo))
	userCommandBus.Register(&userCommand.UpdateUserCommand{}, userHandler.NewUpdateUserHandler(userRepo))

	userQuerydBus.Register(&userQuery.FindUserByIdQuery{}, userHandler.NewFindUserByIdHandler(userRepo))
	userQuerydBus.Register(&userQuery.FindUserByEmailQuery{}, userHandler.NewFindUserByEmailHandler(userRepo))
	userQuerydBus.Register(&userQuery.FindUserByUsernameQuery{}, userHandler.NewFindUserByUsernameHandler(userRepo))
	userQuerydBus.Register(&userQuery.ListUsersByTenantQuery{}, userHandler.NewListUsersByTenantHandler(userRepo))

	// UserUC := usecases.NewUserUseCase(UserRepo)
	// PolicyUC := usecases.NewPolicyUseCase(PolicyRepo)
	// PermissionUC := usecases.NewPermissionUseCase(PermissionRepo)
	// GroupUC := usecases.NewGroupUseCase(GroupRepo)

	app := App{
		db:       db,
		UserRepo: userRepo,
		// PermissionRepo: permissionRepo,
		// PolicyRepo:     policyRepo,
		UserServerSvc:  userServiceSvc,
		UserCommandBus: userCommandBus,
		userQuerydBus:  userQuerydBus,
	}

	return &app, func() {}, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

	for {
		select {
		case <-ctx.Done():
			slog.Info("Shuting Down the client.")
			return
		default:
			slog.Info("default interception ....")
		}
	}
	// forever := struct{}{}
	// <-forever
}
