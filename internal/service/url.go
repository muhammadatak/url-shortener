package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"
	"time"

	"github.com/muhammadatak/url-shortener/domain"
	"github.com/muhammadatak/url-shortener/dto"
)

type urlService struct {
	urlRepository domain.UrlRepository
}

// Create implements [domain.UrlService].

func NewUrl(urlRepository domain.UrlRepository) domain.UrlService {
	return &urlService{
		urlRepository: urlRepository,
	}
}

// Redirect implements [domain.UrlService].
func (u *urlService) Redirect(ctx context.Context, code string) (string, error) {
	url, err := u.urlRepository.FindByCode(ctx, code)

	if err != nil {
		return "", err
	}

	if url.LongUrl == "" {
		return "", errors.New("notfound")
	}

	return url.ShortCode, nil

}

// Shorten implements [domain.UrlService].
func (u *urlService) Shorten(ctx context.Context, data dto.CreateUrlData) (dto.UrlData, error) {

	shortCode, err := u.GenerateCode()
	if err != nil {
		return dto.UrlData{}, err
	}

	url := domain.Url{
		LongUrl:   data.LongUrl,
		ShortCode: shortCode,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	err = u.urlRepository.Save(ctx, &url)

	return dto.UrlData{
		Id:        int(url.Id),
		LongUrl:   url.LongUrl,
		ShortCode: url.ShortCode,
	}, err

}

func (u *urlService) GenerateCode() (string, error) {
	code := make([]byte, 0, domain.CODELENGHT)
	for i := 0; i < domain.CODELENGHT; i++ {
		char, err := rand.Int(rand.Reader, big.NewInt(62))

		if err != nil {
			return "", err
		}
		result := domain.ALPHABET[char.Int64()]
		code = append(code, result)
	}

	return string(code), nil
}
