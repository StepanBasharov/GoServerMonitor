package schemas

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
}

type SuperUserConfig struct {
	Username string
	Password string
	Email    string
}

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Superuser SuperUserConfig
}
