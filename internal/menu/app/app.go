package app

import (
	"context"
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

	app := App{
		DB:     db,
		CmdBus: cmdBus,
		QryBus: qryBus,
	}
	return &app, nil
}

func (app *App) Worder(ctx context.Context, github.com/maruki00/deligo <-chan amqp091.Delivery) {

}
