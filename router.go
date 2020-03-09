package main

import (
	"belajar_api/app"
	student "belajar_api/modules/v1/makul/usecase"

	"github.com/labstack/echo"
)

func initRoute(a app.App, e *echo.Echo) {
	// satu service akan berada dalam 1 grup routing
	// harus di bawah '/public/api/v1'
	crudTemplateGroupingPath := e.Group("/public/api/v1")
	// Route di bawah akan dikelola oleh handler
	crudTemplateGroupingPath.GET("/makuls", student.HandleGetMakul(a))
	crudTemplateGroupingPath.POST("/makuls", student.HandleCreateMakul(a))
	crudTemplateGroupingPath.PUT("/makuls", student.HandleUpdateMakul(a))
	crudTemplateGroupingPath.DELETE("/makuls", student.HandleDeleteMakul(a))

}
