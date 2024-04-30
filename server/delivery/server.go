package delivery

import (
	"fmt"
	"furnishop/server/config"
	"furnishop/server/delivery/middleware"
	"furnishop/server/manager"
	"furnishop/server/util/common"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	return &Server{
		uc:         usecaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
	}
}
