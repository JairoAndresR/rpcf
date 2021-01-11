package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	accountInitializer "rpcf/app/accounts/initializer"
	"rpcf/app/accounts/middleware"
	gruplacInitializer "rpcf/app/gruplacs/initializers"
	productsInitializer "rpcf/app/products/initializers"
	"rpcf/core"
)

type server struct {
	router     *gin.Engine
	authorizer *middleware.Authorizer
}

func NewServer() *server {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use()
	s := &server{router: r}
	s.loadAuthorizer()
	s.setup()
	s.migrations()
	return s
}

func (s *server) setup() {
	s.setupRoutes()
	setupLogger(s)
}

func (s *server) setupRoutes() {
	setupLoginRoutes(s)
	setupScrapingRoutes(s)
	setupProductDefinitionRoutes(s)
	setupGruplacDefinitionRoutes(s)
	setupProductsRoutes(s)
	setupAuthorsRoutes(s)
}

func (s *server) migrations() {
	accountInitializer.Migrate()
	gruplacInitializer.Migrate()
	productsInitializer.Migrate()
}

func (s *server) Run() {
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)
	s.router.Run(address)
}

func (s *server) loadAuthorizer() {
	var authorizer *middleware.Authorizer
	invokeFunc := func(a *middleware.Authorizer) {
		authorizer = a
	}
	err := core.Injector.Invoke(invokeFunc)

	if err != nil {
		panic(err)
	}

	s.authorizer = authorizer
}
