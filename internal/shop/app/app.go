package app

import (
	"context"
	"deligo/internal/shop/app/command"
	"deligo/internal/shop/app/handler"
	"deligo/internal/shop/infra/repository"
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

	cmdBus.Register(&command.CloseShopCommand{}, handler.NewCloseShopHandler())
	cmdBus.Register(&command.AcceptOrderCommand{}, handler.NewAcceptOrderHandler(shopRepo))
	cmdBus.Register(&command.OpenShopCommand{}, handler.NewOpenShopHandler(shopRepo))
	cmdBus.Register(&command.UpdateShopCommand{}, handler.NewUpdateShopHandler(shopRepo))

	app := App{
		DB:     db,
		CmdBus: cmdBus,
		QryBus: qryBus,
	}
	return &app, nil
}

func (app *App) Worder(ctx context.Context, delivery <-chan amqp091.Delivery) {

}
