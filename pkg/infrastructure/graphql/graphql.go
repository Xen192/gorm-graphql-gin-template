package graphql

import (
	"mygpt/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

// NewServer generates graphql server
func NewServer() *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema())
	srv.AddTransport(&transport.Websocket{})
	return srv
}
