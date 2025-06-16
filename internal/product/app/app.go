package app

import (
	"context"

	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
	"github.com/rabbitmq/amqp091-go"
)

type App struct {

	DB *pkgPostgres.DBHandler

	CmdBus *pkgCqrs.CommandBus
	QryBus *pkgCqrs.QueryBus


}

func Init(cgfPath string) *App {
	
	db, err := pkgPostgres.NewDB("")
	if err != nil {
		panic(err)
	}

	cmdBus := pkgCqrs.NewCommandBus()
	qryBus := pkgCqrs.NewQueryBus()


	return &App{
		DB: db,
		CmdBus: cmdBus,
		QryBus: qryBus,
	}
}
func (app *App) Worder(ctx context.Context, github.com/maruki00/deligo <-chan amqp091.Delivery) {

}
