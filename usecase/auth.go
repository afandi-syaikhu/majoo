package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/afandi-syaikhu/majoo/constant"
	"github.com/afandi-syaikhu/majoo/model"
	"github.com/afandi-syaikhu/majoo/repository"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type Auth struct {
	UserRepo repository.UserRepository
	Config   *model.Config
}

func NewAuthUseCase(userRepo repository.UserRepository, config *model.Config) AuthUseCase {
	return &Auth{
		UserRepo: userRepo,
		Config:   config,
	}
}

func (_a *Auth) Login(ctx context.Context, data model.Auth) (*model.Response, error) {
	response := &model.Response{}
	user, err := _a.UserRepo.FindByUsernameAndPassword(ctx, data)
	if err != nil && err == sql.ErrNoRows {
		response.Message = constant.InvalidCredential
		return response, nil
	}

	if err != nil {
		log.Errorf("[%s] => %s", "AuthUC.Login", err.Error())
		return nil, err
	}

	jwtToken, err := generateJWT(user.Username, _a.Config.Jwt.SecretKey, _a.Config.Jwt.ExpTime)
	if err != nil {
		log.Errorf("[%s] => %s", "AuthUC.Login", err.Error())
		return nil, err
	}

	token := &model.Token{
		Token: jwtToken,
		Type:  "Bearer",
	}

	response.Data = token
	response.Success = true

	return response, nil
}

func generateJWT(username, secretKey string, duration int) (string, error) {
	mySigningKey := []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_name"] = username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(duration)).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
