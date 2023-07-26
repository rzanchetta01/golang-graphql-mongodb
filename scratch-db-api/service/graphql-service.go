package service

import (
	"log"
	"scratch-db-api/entities"
	"scratch-db-api/infra/repository"

	"github.com/graphql-go/graphql"
)

func Query() graphql.Schema {

	log.Println("querry called")

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"getAllRecords": &graphql.Field{
				Type: graphql.NewList(entities.ModelType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return repository.GetAll(), nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query: rootQuery,
	}

	log.Println(schemaConfig)

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("ERROR RETRIEVING GRAPH QUERY -> ", err)
	}

	return schema
}
