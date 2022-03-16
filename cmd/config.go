package main

import (
	"log"
	"os"

	"github.com/golobby/dotenv"
)

type cfg struct {
	Database struct {
		User        string `env:"USER"`
		Password    string `env:"PASSWORD"`
		Host        string `env:"HOST"`
		Db          string `env:"DB"`
		MaxOpenConn int    `env:"MAX_OPEN_CON"`
		MaxIdleConn int    `env:"MAX_IDLE_CON"`
		MaxTTLConn  int    `env:"MAX_TTL_CONN"`
	}
}

func initConfig() (*cfg, error) {
	/* os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("DB"), os.Getenv("MAXOPENCONN"), os.Getenv("MAXIDLECONN"), os.Getenv("MAXTTLCONN") */
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := &cfg{}

	err = dotenv.NewDecoder(file).Decode(&cfg)
	if err == nil {
		return nil, err
	}
	return cfg, nil
}
