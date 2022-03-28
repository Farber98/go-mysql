package main

import (
	"go-mysql/internal/controllers"
	"go-mysql/internal/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initRouter(db *db.MySQLHandler) *echo.Echo {
	r := echo.New()
	initRoutes(r, db)

	/* MIDDLEWARES */
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
	return r
}

func initRoutes(r *echo.Echo, db *db.MySQLHandler) {

	r.GET("/test", controllers.Prueba)

	cuadrosController := &controllers.CuadrosController{Db: db}
	r.GET("/cuadros/estado/:id", cuadrosController.EstadoCuadros)
}
