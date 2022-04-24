package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"
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

	jwtToken, err := _a.generateJWT(user.Username)
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

func (_a *Auth) ValidateToken(ctx context.Context, headerAuth string) (*model.TokenExtraction, error) {
	errUnauthorized := errors.New(constant.Unauthorized)
	errInternalServer := errors.New(constant.InternalServerError)
	errExpiredToken := errors.New(constant.ExpiredToken)
	tokens := strings.Split(headerAuth, " ")
	if len(tokens) != 2 {
		return nil, errUnauthorized
	}

	mySigningKey := []byte(_a.Config.Jwt.SecretKey)
	token, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInternalServer
		}

		return mySigningKey, nil
	})

	if err != nil {
		return nil, errUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errUnauthorized
	}

	exp, ok := claims[_a.Config.Jwt.Content.ExpTime].(float64)
	if !ok {
		return nil, errInternalServer
	}

	expToken := int64(exp)
	if expToken < time.Now().Unix() {
		return nil, errExpiredToken
	}

	username, ok := claims[_a.Config.Jwt.Content.Username].(string)
	if !ok {
		return nil, errInternalServer
	}

	return &model.TokenExtraction{
		ExpTime:  expToken,
		Username: username,
	}, nil
}

func (_a *Auth) generateJWT(username string) (string, error) {
	mySigningKey := []byte(_a.Config.Jwt.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[_a.Config.Jwt.Content.Username] = username
	claims[_a.Config.Jwt.Content.ExpTime] = time.Now().Add(time.Minute * time.Duration(_a.Config.Jwt.ExpTime)).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
