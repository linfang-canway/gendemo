package main

import (
	"gendemo/dal/model"

	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery, /*WithQueryInterface, WithoutContext*/

		WithUnitTest: true,
	})

	g.ApplyBasic(model.User{}) // Associations

	g.Execute()
}
