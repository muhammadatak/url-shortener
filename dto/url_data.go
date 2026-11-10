package dto

type UrlData struct {
	Id        int    `json:"id"`
	ShortCode string `json:"short_code"`
	LongUrl   string `json:"long_url"`
	Count     int    `json:"count"`
}

type CreateUrlData struct {
	LongUrl string `json:"long_url"`
}
