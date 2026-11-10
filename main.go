package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/url-shortener/internal/api"
	"github.com/muhammadatak/url-shortener/internal/config"
	"github.com/muhammadatak/url-shortener/internal/connection"
	"github.com/muhammadatak/url-shortener/internal/repository"
	"github.com/muhammadatak/url-shortener/internal/service"
)

func main() {
	cnf := config.Get()
	dbConn := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	urlRepository := repository.NewUrl(dbConn)
	urlService := service.NewUrl(urlRepository)

	api.NewApi(app, urlService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}
