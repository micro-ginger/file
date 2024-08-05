package app

import "github.com/ginger-repository/sql"

func (a *App[acc, f]) initializeDatabases() {
	a.initializeSql()
}

func (a *App[acc, f]) initializeSql() {
	a.Sql = sql.New(
		a.Logger.WithTrace("sql"),
		a.Registry.ValueOf("sql"),
	)
	if err := a.Sql.Initialize(); err != nil {
		panic(err)
	}
}
