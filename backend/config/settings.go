package config

// Settings exported
type Settings struct {
	Server   Server
	Database Database
	Redis    Redis
}

// Server exported
type Server struct {
	Port int
}

// Database exported
type Database struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	Charset  string
}

// Redis redis configration
type Redis struct {
	Host string
	Port int
	DB   int
}
