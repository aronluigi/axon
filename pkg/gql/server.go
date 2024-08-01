package gql

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aronluigi/axon/pkg/gql/graph"
	"github.com/aronluigi/axon/pkg/gql/graph/generated"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func StartServer(port int) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	resolver := graph.NewResolver()
	config := generated.Config{Resolvers: resolver}

	srv := handler.New(generated.NewExecutableSchema(config))

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", c.Handler(srv))

	fmt.Printf("GraphQL Server http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
