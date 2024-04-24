package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natchapol2347/woodman_web_back/adaptor/storage"
	"github.com/natchapol2347/woodman_web_back/handler"
	"github.com/natchapol2347/woodman_web_back/internal/config"
	"github.com/natchapol2347/woodman_web_back/internal/database"
	"github.com/natchapol2347/woodman_web_back/service"
)

func main() {
	// time.DateOnly
	fmt.Println(config.LoadConfig())
	db, _ := database.NewPostgres(config.LoadConfig())
	storageClient := storage.NewStorage(db)

	s := service.NewService(storageClient)
	handlerPortfolio := handler.NewPortfolioHandler(s)
	handlerJob := handler.NewJobHandler(s)
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/project/:id", handlerPortfolio.GetProject)
	e.GET("/projects", handlerPortfolio.GetManyProjects)
	e.POST("/project", handlerPortfolio.PostProject)
	e.DELETE("/remove-project/:id", handlerPortfolio.DeleteProject)
	e.PUT("/update-project/:id", handlerPortfolio.UpdateProject)
	e.GET("/jobs", handlerJob.GetManyJobs)
	e.GET("/job/:id", handlerJob.GetJob)
	e.POST("/job", handlerJob.PostJob)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))

}
