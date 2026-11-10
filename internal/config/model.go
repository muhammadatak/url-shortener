package config

type Config struct {
	Server   Server
	Database Database
}

type Database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
	Tz   string
}

type Server struct {
	Host string
	Port string
}
