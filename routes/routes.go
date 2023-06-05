package routes

import (
	"go-echo/db"
	"go-echo/handler"
	"go-echo/repository"
	"go-echo/services"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *echo.Echo {
	e := echo.New()

	biodataRepository := repository.NewRepository(db.Init())
	biodataService := services.NewService(biodataRepository)
	biodataHandler := handler.NewBiodataHandler(biodataService)

	e.GET("/", handler.Root)
	e.GET("/biodata", biodataHandler.GetAll)
	e.GET("/biodata/:id", biodataHandler.FindByID)
	e.POST("/biodata", biodataHandler.Create)
	e.PUT("/biodata/:id", biodataHandler.Update)
	e.DELETE("/biodata/:id", biodataHandler.Delete)
	return e
}
