package main

import (
	"awesomeProject/handler"
	"awesomeProject/model"
	"awesomeProject/response"
	"awesomeProject/router"
	"awesomeProject/util"
	"html/template"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

func main() {
	//config

	api := echo.New()
	api.HideBanner = true
	api.HTTPErrorHandler = response.HTTPErrorHandler

	renderer := &util.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("templates/admin/*.html")),
	}
	api.Renderer = renderer

	conn, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err) //todo change logger
	}

	dbModel := model.New(conn, "config")
	// Middleware
	api.Use(middleware.Recover())
	api.Use(middleware.Logger())
	api.Use(middleware.RequestID())

	// routing
	router.Admin(api, handler.NewAdmin(dbModel))

	//Start Echo Web Server
	api.Logger.Fatal(api.Start("localhost:5000"))

}
