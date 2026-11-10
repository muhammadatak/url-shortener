package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/muhammadatak/url-shortener/internal/config"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open connection", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping", err.Error())
	}

	return db
}
