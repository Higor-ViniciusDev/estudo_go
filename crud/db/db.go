package db

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("sqlserver", "sqlserver://sa:Negrao3d@DESKTOP-A3E2KM7?database=estudo&connection+timeout=30")

	if err != nil {
		panic(err.Error())
	}

	return db
}
