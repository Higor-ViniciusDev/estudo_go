package database

import (
	"estudo_go/GIN/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("ERRO AO CONECTAR COM O BANCO DE DADOS")
	}

	DB.AutoMigrate(&models.Aluno{})
}
