package main

import (
	"database/sql"
	"fmt"

	"github.com/afandi-syaikhu/majoo/config"
	"github.com/afandi-syaikhu/majoo/delivery/rest"
	"github.com/afandi-syaikhu/majoo/repository"
	"github.com/afandi-syaikhu/majoo/usecase"
	"github.com/labstack/echo/v4"
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

	// init repo
	userRepo := repository.NewUserRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	outletRepo := repository.NewOutletRepository(db)

	// init usecase
	authUC := usecase.NewAuthUseCase(userRepo, cfg)
	merchantUC := usecase.NewMerchantUseCase(merchantRepo)
	outletUC := usecase.NewOutletUseCase(outletRepo)

	// init echo framework
	e := echo.New()

	// init handler
	rest.NewAuthHandler(e, authUC)
	rest.NewMerchantHandler(e, authUC, merchantUC)
	rest.NewOutletHandler(e, authUC, outletUC)

	e.Logger.Fatal(e.Start(":8080"))
}
