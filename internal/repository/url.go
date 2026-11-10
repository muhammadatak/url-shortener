package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/url-shortener/domain"
)

type urlRepository struct {
	db *goqu.Database
}

func NewUrl(con *sql.DB) domain.UrlRepository {
	return &urlRepository{
		db: goqu.New("default", con),
	}
}

func (u *urlRepository) FindByCode(ctx context.Context, code string) (result domain.Url, err error) {
	dataset := u.db.From("urls").Where(goqu.C("short_code").Eq("code"))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements [domain.UrlRepository].
func (u *urlRepository) Save(ctx context.Context, url *domain.Url) error {
	executor := u.db.Insert("urls").Rows(url).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
