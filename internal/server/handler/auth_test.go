package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/handler"
	mock "github.com/pisarevaa/gophkeeper/internal/server/mocks"
	"github.com/pisarevaa/gophkeeper/internal/server/router"
	"github.com/pisarevaa/gophkeeper/internal/server/service/auth"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
)

type ServerTestSuite struct {
	suite.Suite
	config config.Config
	client *resty.Client
	token  string
}

const (
	userID   = 1
	email    = "test@example.com"
	password = "12345678"
)

func (suite *ServerTestSuite) SetupSuite() {
	suite.config = config.NewConfig()
	suite.client = resty.New()
	token, err := utils.GenerateJWTString(suite.config.Security.TokenExpSec, suite.config.Security.SecretKey, userID)
	suite.Require().NoError(err)
	suite.token = token
}

func TestAgentSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (suite *ServerTestSuite) TestRegisterUser() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	m := mock.NewMockAuthStorage(ctrl)
	m.EXPECT().
		GetUserByEmail(gomock.Any(), gomock.Any()).
		Return(model.User{}, errors.New("email не найден"))
	m.EXPECT().
		RegisterUser(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(model.User{
			ID:        1,
			Email:     email,
			Password:  password,
			CreatedAt: time.Now(),
		}, nil)

	authService := auth.NewService(
		auth.WithConfig(suite.config),
		auth.WithStorage(m),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithAuthService(authService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	user := model.RegisterUser{
		Email:    email,
		Password: password,
	}
	resp, err := suite.client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(ts.URL + "/auth/register")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
}

func (suite *ServerTestSuite) TestLogin() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	m := mock.NewMockAuthStorage(ctrl)
	passwordHash, _ := utils.GetPasswordHash(password, suite.config.Security.SecretKey)
	m.EXPECT().
		GetUserByEmail(gomock.Any(), gomock.Any()).
		Return(model.User{
			ID:        1,
			Email:     email,
			Password:  passwordHash,
			CreatedAt: time.Now(),
		}, nil)

	authService := auth.NewService(
		auth.WithConfig(suite.config),
		auth.WithStorage(m),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithAuthService(authService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	user := model.RegisterUser{
		Email:    email,
		Password: password,
	}

	resp, err := suite.client.R().
		SetBody(user).
		SetHeader("Content-Type", "application/json").
		Post(ts.URL + "/auth/login")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
}
