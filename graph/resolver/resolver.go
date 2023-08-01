package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"mygpt/graph"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// NewSchema creates NewExecutableSchema
func NewSchema() graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: &Resolver{},
	})
}
