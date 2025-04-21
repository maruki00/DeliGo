package app

import (
	"context"
	"deligo/cmd/user/configs"
	"deligo/internal/iam/app/usecases"
	userCommands "deligo/internal/iam/app/user/commands"
	userHandlers "deligo/internal/iam/app/user/handlers"
	userQueries "deligo/internal/iam/app/user/queries"
	"deligo/internal/iam/domain/contracts"
	"deligo/internal/iam/infra/repositories"
	pkgCqrs "deligo/pkg/cqrs"
	pkgPostgres "deligo/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db             *pkgPostgres.PGHandler
	UserRepo       contracts.IUserRepository
	GroupRepo      contracts.IGroupRepository
	PermissionRepo contracts.IPermissionRepository
	PolicyRepo     contracts.IPolicyRepository
	UserUC         *usecases.UserUseCase
	PolicyUC       *usecases.PolicyUseCase
	PermissionUC   *usecases.PermissionUseCase
	GroupUC        *usecases.GroupUseCase
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *configs.Config) (*App, func(), error) {

	fmt.Println("dsn : ", cfg.Postgres.Dsn)
	db, err := pkgPostgres.NewDB(cfg.Postgres.Dsn)
	if err != nil {
		return nil, func() {}, err
	}
	userCommandBus := pkgCqrs.NewCommandBus()
	userQuerydBus := pkgCqrs.NewQueryBus()

	UserRepo := repositories.NewUserRepository(db)
	GroupRepo := repositories.NewGroupRepository()
	PermissionRepo := repositories.NewPermissionRepository()
	PolicyRepo := repositories.NewPolicyRepository()

	userCommandBus.Register(&userCommands.CreateUserCommand{}, userHandlers.NewCreateUserHandler(UserRepo))
	userCommandBus.Register(&userCommands.DeleteUserCommand{}, userHandlers.NewDeleteUserHandler(UserRepo))
	userCommandBus.Register(&userCommands.UpdateUserCommand{}, userHandlers.NewUpdateUserHandler(UserRepo))

	userQuerydBus.Register(&userQueries.FindUserByIdQuery{}, userHandlers.NewFindUserByIdHandler(UserRepo))
	userQuerydBus.Register(&userQueries.FindUserByEmailQuery{}, userHandlers.NewDeleteUserHandler(UserRepo))
	userQuerydBus.Register(&userQueries.FindUserByEmailQuery{}, userHandlers.NewUpdateUserHandler(UserRepo))

	// UserUC := usecases.NewUserUseCase(UserRepo)
	// PolicyUC := usecases.NewPolicyUseCase(PolicyRepo)
	// PermissionUC := usecases.NewPermissionUseCase(PermissionRepo)
	// GroupUC := usecases.NewGroupUseCase(GroupRepo)

	app := &App{
		db:             db,
		UserRepo:       UserRepo,
		GroupRepo:      GroupRepo,
		PermissionRepo: PermissionRepo,
		PolicyRepo:     PolicyRepo,
		UserUC:         UserUC,
		PolicyUC:       PolicyUC,
		PermissionUC:   PermissionUC,
		GroupUC:        GroupUC,
	}

	return app, func() {}, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

	for {
		select {
		case <-ctx.Done():
			slog.Info("Shuting Down the client.")
			break
		default:
			slog.Info("default interception ....")
		}
	}
	// forever := struct{}{}
	// <-forever
}
