package app

import (
	"context"
	"deligo/internal/shop/app/commands"
	"deligo/internal/shop/app/handlers"
	repository "deligo/internal/shop/infra/respositories"
	pkgCqrs "deligo/pkg/cqrs"
	pkgPostgres "deligo/pkg/postgres"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	DB     *pkgPostgres.PGHandler
	CmdBus *pkgCqrs.CommandBus
	QryBus *pkgCqrs.QueryBus
}

func NewApp(cfg any) (*App, error) {

	db, err := pkgPostgres.NewDB("")
	if err != nil {
		return nil, err
	}

	cmdBus := pkgCqrs.NewCommandBus()
	qryBus := pkgCqrs.NewQueryBus()

	shopRepo := repository.NewShopRepository(db)

	cmdBus.Register(&commands.CloseShopCommand{}, handlers.NewCloseShopHandler(shopRepo))
	cmdBus.Register(&commands.CreateShopCommand{}, handlers.NewCreateShopHandler(shopRepo))

	app := App{
		DB:     db,
		CmdBus: cmdBus,
		QryBus: qryBus,
	}
	return &app, nil
}

func (app *App) Worder(ctx context.Context, delivery <-chan amqp091.Delivery) {

}
