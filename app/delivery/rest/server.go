package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	accountInitializer "rpcf/app/accounts/initializer"
	gruplacInitializer "rpcf/app/gruplacs/initializers"
)

type server struct {
	router *gin.Engine
}

func NewServer() *server {
	r := gin.Default()
	r.Use(CORSMiddleware())
	s := &server{router: r}
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
}

func (s *server) migrations() {
	accountInitializer.Migrate()
	gruplacInitializer.Migrate()
}

func (s *server) Run() {
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf(":%s", port)
	s.router.Run(address)
}
