package handler

import (
	"os"

	"github.com/khdiyz/book-store/pkg/service"
	"github.com/khdiyz/book-store/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	services *service.Service
	log      *logger.Logger
}

func NewHandler(services *service.Service, log *logger.Logger) *Handler {
	return &Handler{services: services, log: log}
}

func (h *Handler) InitRoutes() *echo.Echo {
	app := echo.New()

	loggerConfig := middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human}\n",
		Output: os.Stdout,
	}

	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(loggerConfig))

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := app.Group("/api/v1")
	{
		category := v1.Group("/categories")
		{
			category.POST("", h.createCategory)
			category.GET("", h.getCategoryList)
			category.GET("/:id", h.getCategoryByID)
			category.PUT("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}
	}

	return app
}
