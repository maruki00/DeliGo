package app

import (
	"context"
	"deligo/cmd/iam/configs"
	userServerServices "deligo/internal/iam/app/user"
	userCommands "deligo/internal/iam/app/user/commands"
	userHandlers "deligo/internal/iam/app/user/handlers"
	userQueries "deligo/internal/iam/app/user/queries"
	"deligo/internal/iam/domain/contracts"
	"deligo/internal/iam/infra/repositories"
	pkgCqrs "deligo/pkg/cqrs"
	pkgPostgres "deligo/pkg/postgres"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db             *pkgPostgres.PGHandler
	UserRepo       contracts.IUserRepository
	PermissionRepo contracts.IPermissionRepository
	PolicyRepo     contracts.IPolicyRepository
	UserServerSvc  *userServerServices.UserServerService
	UserCommandBus *pkgCqrs.CommandBus
	UserQuerydBus  *pkgCqrs.QueryBus

	// UserUC         *usecases.UserUseCase
	// Polshared_valueobject.ID	// PermissionUC   *usecases.PermissionUseCase
	// GroupUC        *usecases.GroupUseCase
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

	userRepo := repositories.NewUserRepository(db)

	//permissionRepo := repositories.NewPermissionRepository()

	userServiceSvc := userServerServices.NewUserUseCase(userCommandBus, userQuerydBus)

	userCommandBus.Register(&userCommands.CreateUserCommand{}, userHandlers.NewCreateUserHandler(userRepo))
	userCommandBus.Register(&userCommands.DeleteUserCommand{}, userHandlers.NewDeleteUserHandler(userRepo))
	userCommandBus.Register(&userCommands.UpdateUserCommand{}, userHandlers.NewUpdateUserHandler(userRepo))

	userQuerydBus.Register(&userQueries.FindUserByIdQuery{}, userHandlers.NewFindUserByIdHandler(userRepo))
	userQuerydBus.Register(&userQueries.FindUserByEmailQuery{}, userHandlers.NewFindUserByEmailHandler(userRepo))
	userQuerydBus.Register(&userQueries.FindUserByUsernameQuery{}, userHandlers.NewFindUserByUsernameHandler(userRepo))
	userQuerydBus.Register(&userQueries.ListUsersByTenantQuery{}, userHandlers.NewListUsersByTenantHandler(userRepo))

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
		UserQuerydBus:  userQuerydBus,
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
