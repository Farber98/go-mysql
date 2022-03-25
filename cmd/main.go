package main

import (
	"fmt"
	"go-mysql/internal/db"
	"log"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	configPath := "../cfg.toml"
	conf, err := initConfig(configPath)
	if err != nil {
		log.Panic("Config file error.")
	}
	db, err := db.InitDb(conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Db, conf.Database.MaxOpenConn, conf.Database.MaxIdleConn, conf.Database.MaxTtlConn)
	if err != nil {
		log.Panic("Initializing db error.")
	}
	defer db.Conn.Close()
	if err := db.Conn.Ping(); err != nil {
		log.Println("Ping error.")
	}
	fmt.Println(conf)
	r := initRouter(db)
	r.Logger.Fatal(r.Start(conf.Server.Listen))

}
