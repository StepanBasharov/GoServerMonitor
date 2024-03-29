package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"servermonitor/pkg/db"
	"servermonitor/pkg/routers"
	"servermonitor/pkg/types"
)

func setUpLoggers(server *echo.Echo) {
	// set up logging middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} - ${uri} - ${status}\n",
	}))

	// set up go logger
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func ReadFlagConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.toml", "Config File Path")
	flag.Parse()
	log.Println("Load CLI args")
	return configPath
}

func LoadConfig() *types.Config {
	configPath := ReadFlagConfigPath()
	var conf types.Config
	errConf := conf.LoadConfig(configPath)
	if errConf != nil {
		log.Panicln("Cannot load config")
	}
	log.Println("Load Config File")
	return &conf
}

func SetUpDatabase(conf *types.Config) *gorm.DB {
	conn, errDb := db.CreateConnection(conf.Database)
	if errDb != nil {
		log.Panicln("Cannot connect to database")
	}
	log.Println("Connect to Database")
	return conn
}

func CreateSuperUser(conf *types.Config, conn *gorm.DB) {
	db.CreateSuperUser(conn, conf.Superuser)
	log.Println("Super User created")
}

func SetUpRouters(server *echo.Echo, conn *gorm.DB) {
	api := server.Group("/api")
	routers.RegisterServiceRouters(api)
	routers.RegisterHardwareMonitorRouters(api, conn)
	routers.RegisterUserRouters(api, conn)
	log.Println("Register routers")
}

func main() {
	server := echo.New()
	conf := LoadConfig()
	setUpLoggers(server)
	conn := SetUpDatabase(conf)
	CreateSuperUser(conf, conn)
	SetUpRouters(server, conn)
	server.Logger.Info(server.Start(conf.Server.GetBuildAddress()))
}
