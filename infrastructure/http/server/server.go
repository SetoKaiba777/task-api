package server

import (
	"go-challenger/adapter/repository"
	coreRepository "go-challenger/core/repository"
	appConfig "go-challenger/infrastructure/config"
	"go-challenger/infrastructure/database"
	"go-challenger/infrastructure/http/router"
	"go-challenger/infrastructure/logger"
	"strconv"
	"time"
)

type config struct {
	
	configApp  *appConfig.AppConfig
	webServer  router.Server
	db         *database.MySQLConnection
	repo       coreRepository.TaskRepository
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WithAppConfig() *config {
	var err error
	c.configApp, err = appConfig.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *config) InitLogger() *config {
	logger.NewZapLogger()
	logger.Infof("Log has been successfully configured")
	return c
}

func (c *config) WithDB() *config {
	db, err := database.NewMySQLConnection(c.configApp.MySQL.Host)
	if err != nil {
		logger.Fatal(err)
	}

	c.db = db
	logger.Infof("DB has been successfully configured")
	return c
}


func (c *config) WithRepository() *config {
	c.repo = repository.NewTaskRepository(c.db)
	logger.Infof("Repository has been successfully configured")
	return c
}

func (c *config) WithWebServer() *config {
	intPort, err := strconv.ParseInt(c.configApp.Application.Server.Port, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}

	intDuration, err := time.ParseDuration(c.configApp.Application.Server.Timeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	c.webServer = router.NewGinServer(intPort, intDuration, c.repo)
	logger.Infof("Web server has been successfully configurated")
	return c
}

func (c *config) Start() {
	logger.Infof("App running on port %s", c.configApp.Application.Server.Port)
	c.webServer.Listen()
	
}