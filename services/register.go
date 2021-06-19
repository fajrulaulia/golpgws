package golpgws

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	app "github.com/fajrulaulia/golpgws/app"
)

func Register(router *mux.Router, db *pgx.Conn) {
	ctx := context.Background()
	UserService := &app.UserService{Name: "User Service", Router: router, Context: ctx, Db: db}
	UserService.Register()
}
