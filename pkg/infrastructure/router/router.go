package router

import (
	"mygpt/pkg/infrastructure/controller"
	"mygpt/pkg/infrastructure/datastore"
	"mygpt/pkg/infrastructure/middleware"
	"mygpt/pkg/infrastructure/repository"
	"mygpt/pkg/infrastructure/service"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	TestPath       = "/test"
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
	FilePath       = "/file"
)

// New creates route endpoint
func New(srv *handler.Server) *gin.Engine {
	switch os.Getenv("APP_ENV") {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	f := g.Group(FilePath)
	{
		fr := repository.NewFileUploadRepository(datastore.GetInstance())
		fs := service.NewFileUploadService(fr, time.Second*10)
		fc := controller.FileUploadController{Service: fs}

		f.PUT("/:parent_id", fc.Put)
		f.GET("/get_url/:file_id", fc.GenerateURL)
		f.GET("/:file_id", fc.Get)
		f.DELETE("/:file_id", fc.Delete)
	}

	if os.Getenv("APP_ENV") != "dev" {
		g.Use(middleware.AuthMiddleware())
	}
	{
		g.POST(QueryPath, func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		})
		g.GET(PlaygroundPath, func(c *gin.Context) {
			playground.Handler("GraphQL", QueryPath).ServeHTTP(c.Writer, c.Request)
		})
	}

	return g
}
