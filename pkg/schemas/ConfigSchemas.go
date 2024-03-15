package schemas

type ServerAuth struct {
	Secret string
}

type ServerConfig struct {
	Host string
	Port int
	Auth ServerAuth
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

type RedisConfig struct {
	Host string
	Port int
}

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Superuser SuperUserConfig
	Redis     RedisConfig
}
