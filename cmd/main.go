package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"servermonitor/pkg/db"
	"servermonitor/pkg/routers"
	"servermonitor/pkg/tools"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	server := echo.New()
	var configPath string
	flag.StringVar(&configPath, "config", "./config.toml", "Config File Path")
	flag.Parse()
	log.Println("Load CLI args")
	conf, errConf := tools.ConfigReader(configPath)
	if errConf != nil {
		log.Panicln("Cannot load config")
	}
	log.Println("Load Config File")
	conn, errDb := db.CreateConnection(conf.Database)
	if errDb != nil {
		log.Panicln("Cannot connect to database")
	}
	log.Println("Connect to Database")
	db.CreateSuperUser(conn, conf.Superuser)
	log.Println("Super User created")
	api := server.Group("/api")
	routers.RegisterServiceRouters(api)
	routers.RegisterHardwareMonitorRouters(api, conn)
	routers.RegisterUserRouters(api, conn)
	log.Println("Register routers")
	server.Logger.Info(server.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
