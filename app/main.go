package main

import (
	"database/sql"
	"fmt"

	"github.com/afandi-syaikhu/majoo/config"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {

	// read config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
