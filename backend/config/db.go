package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB(path string) *gorm.DB{
	db,err := gorm.Open(sqlite.Open(path),&gorm.Config{})

	if err != nil {
		panic(fmt.Sprintln("Ocorreu um erro ao se conectar ao banco de dados.Erro:",err.Error()))
	}

	fmt.Println("Conectado com sucesso!")
	return db
}

func MigrateAllTables(db *gorm.DB, tables ...interface{}){
	for _,v := range tables {
		db.AutoMigrate(v)
	}
}