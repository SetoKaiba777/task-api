package router

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-challenger/adapter/api/controller"
	corepository "go-challenger/core/repository"
	"go-challenger/core/usecase"
	"go-challenger/infrastructure/logger"

	"github.com/gin-gonic/gin"
)



type (
	Port int64

	Server interface {
		Listen()
	}

	ginEngine struct {
		router     *gin.Engine
		port       int64
		ctx        context.Context
		ctxTimeout time.Duration
		repo       corepository.TaskRepository
	}
)

func NewGinServer(port int64,
	timeout time.Duration,
	repo corepository.TaskRepository) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: timeout,
		repo:       repo,
	}
}

func (engine ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	engine.setAppHandlers(engine.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", engine.port),
		Handler:      engine.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal("Error to starting HTTP Server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(engine.ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Sever Shutdown Failed")
	}
}

func (engine ginEngine) setAppHandlers(router *gin.Engine) {
	router.POST("/v1/tasks", engine.HandlePostTask())
	router.POST("/v1/tasks/list", engine.HandlePostAllTask())
	router.GET("/v1/tasks/:taskId", engine.HandleGetTask())
	router.DELETE("/v1/tasks/:taskId", engine.HandleDeleteTask())
	router.PUT("/v1/tasks/:taskId", engine.HandleUpdateTask())
}

func (e ginEngine) HandlePostTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uc:=usecase.NewSaveUseCase(e.repo)
		c := controller.NewCreateTaskController(uc)
		c.Execute(ctx.Writer, ctx.Request)

	}
}

func (e ginEngine) HandlePostAllTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uc:=usecase.NewSaveAllUseCase(e.repo)
		c := controller.NewCreateAllTaskController(uc)
		c.Execute(ctx.Writer, ctx.Request)

	}
}

func (e ginEngine) HandleGetTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("taskId", ctx.Param("taskId"))
		ctx.Request.URL.RawQuery = query.Encode()

		uc := usecase.NewFindByIdUseCase(e.repo)
		c := controller.NewGetTaskController(uc)
		c.Execute(ctx.Writer, ctx.Request)
	}
}

func (e ginEngine) HandleDeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("taskId", ctx.Param("taskId"))
		ctx.Request.URL.RawQuery = query.Encode()

		uc := usecase.NewDeleteByIdUseCase(e.repo)
		c := controller.NewDeleteTaskController(uc)
		c.Execute(ctx.Writer, ctx.Request)
	}
}

func (e ginEngine) HandleUpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("taskId", ctx.Param("taskId"))
		ctx.Request.URL.RawQuery = query.Encode()

		uc := usecase.NewUpdateUseCase(e.repo)
		c := controller.NewUpdateTaskController(uc)
		c.Execute(ctx.Writer, ctx.Request)
	}
}