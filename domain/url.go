package domain

import (
	"context"
	"database/sql"

	"github.com/muhammadatak/url-shortener/dto"
)

const (
	ALPHABET   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CODELENGHT = 7
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
	Shorten(ctx context.Context, url dto.CreateUrlData) (dto.UrlData, error)
	Redirect(ctx context.Context, code string) (string, error)
}
