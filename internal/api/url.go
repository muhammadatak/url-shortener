package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/url-shortener/domain"
	"github.com/muhammadatak/url-shortener/dto"
)

type urlApi struct {
	urlService domain.UrlService
}

func NewApi(app *fiber.App, urlService domain.UrlService) {
	ua := urlApi{
		urlService: urlService,
	}

	app.Post("/shorten", ua.Shorten)
}

func (ua urlApi) Shorten(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateUrlData

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	res, err := ua.urlService.Shorten(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))

}
