package domain

import (
	"context"
	"database/sql"

	"github.com/muhammadatak/url-shortener/dto"
)

type Url struct {
	Id        int64        `db:"id"`
	ShortCode string       `db:"short_code"`
	LongUrl   string       `db:"long_url"`
	Count     int          `db:"count"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type UrlRepository interface {
	Save(ctx context.Context, url *Url) error
	FindByCode(ctx context.Context, code string) (Url, error)
}

type UrlService interface {
	Shorten(ctx context.Context, url dto.CreateUrlData) error
	Redirect(ctx context.Context, code string) (dto.UrlData, error)
}
