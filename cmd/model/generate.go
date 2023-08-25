package main

import (
	"mygpt/pkg/infrastructure/datastore"

	"github.com/joho/godotenv"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:          "query",
		ModelPkgPath:     "model",
		FieldNullable:    true,
		FieldSignable:    true,
		FieldCoverable:   true,
		FieldWithTypeTag: true,
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		"user_state": func(columnType gorm.ColumnType) (dataType string) {
			return "model_struct.UserStatus"
		},
	}

	g.WithDataTypeMap(dataMap)

	godotenv.Load()
	g.UseDB(datastore.GetInstance())

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	g.Execute()
}
