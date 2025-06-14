package app

import (
	"context"
	"github.com/maruki00/deligo/internal/shop/app/command"
	"github.com/maruki00/deligo/internal/shop/app/handler"
	"github.com/maruki00/deligo/internal/shop/infra/repository"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"

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

	cmdBus.Register(&command.CloseShopCommand{}, handler.NewCloseShopHandler(shopRepo))
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
