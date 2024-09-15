package handler_test

import (
	"net/http/httptest"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/pisarevaa/gophkeeper/internal/server/handler"
	mock "github.com/pisarevaa/gophkeeper/internal/server/mocks"
	"github.com/pisarevaa/gophkeeper/internal/server/model"
	"github.com/pisarevaa/gophkeeper/internal/server/router"
	"github.com/pisarevaa/gophkeeper/internal/server/service/keeper"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

func (suite *ServerTestSuite) TestGetData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		GetDataByUserID(gomock.Any(), gomock.Any()).
		Return(
			[]model.Keeper{
				model.Keeper{
					ID:        1,
					Name:      "Text",
					Data:      "some data",
					Type:      model.TextType,
					UserID:    userID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				model.Keeper{
					ID:        2,
					Name:      "Binary",
					Data:      "url",
					Type:      model.BinaryType,
					UserID:    userID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			}, nil)

	mockMinio.EXPECT().
		GetMany(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			[]string{"url"}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(utils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result []model.DataResponse

	resp, err := suite.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Get(ts.URL + "/api/data")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Len(result, 2)
}

func (suite *ServerTestSuite) TestGetTextDataByID() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "Text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(utils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Get(ts.URL + "/api/data/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.TextType)
}

func (suite *ServerTestSuite) TestGetBinaryDataByID() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "Binary",
				Data:      "url",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	mockMinio.EXPECT().
		GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			"url", nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(utils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Get(ts.URL + "/api/data/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.BinaryType)
}
