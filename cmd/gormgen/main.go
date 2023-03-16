package main

import (
	"gorm.io/gen"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

type Querier interface {
	// SELECT * FROM @@table
	//  {{where}}
	//      user_id=@user_id
	//  {{end}}
	GetByUserID(user_id int) (gen.T, error) // returns struct and error
	// SELECT * FROM @@table
	//  {{where}}
	//      id=@id
	//  {{end}}
	GetByID(id int) (gen.T, error) // returns struct and error
	// SELECT * FROM @@table
	//  {{where}}
	//      id IN (@ids)
	//  {{end}}
	GetByIDs(ids []int) ([]gen.T, error) // returns slice of struct and error
}

func main() {
	generateQueryFiles()
}

func generateQueryFiles() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../../internal/generated/dal",
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	var models = proto.GetDatabaseSchemas()

	// generate from struct in project
	g.ApplyBasic(models...)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, models...)

	g.Execute()
}
