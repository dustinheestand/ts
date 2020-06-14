package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.tmpl.html")),
	}
	e.Renderer = t
	router.LoadHTMLGlob("templates/*.tmpl.html")

	e.GET("/", func(ctx echo.Context) error {
		return c.Render(http.StatusOK, "")
		ctx.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
