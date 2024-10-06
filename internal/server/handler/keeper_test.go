package handler_test

import (
	"net/http/httptest"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/pisarevaa/gophkeeper/internal/server/handler"
	mock "github.com/pisarevaa/gophkeeper/internal/server/mocks"
	"github.com/pisarevaa/gophkeeper/internal/server/router"
	"github.com/pisarevaa/gophkeeper/internal/server/service/keeper"
	"github.com/pisarevaa/gophkeeper/internal/shared/model"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
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
				{
					ID:        1,
					Name:      "text",
					Data:      "some data",
					Type:      model.TextType,
					UserID:    userID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					ID:        2,
					Name:      "Binary",
					Data:      "url",
					Type:      model.BinaryType,
					UserID:    userID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
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
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponseShort

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
		handler.WithValidator(sharedUtils.NewValidator()),
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

func (suite *ServerTestSuite) TestAddTextData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockDB.EXPECT().
		AddData(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse
	body := model.AddTextData{
		Name: "text",
		Data: "some data",
	}

	resp, err := suite.client.R().
		SetBody(body).
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Post(ts.URL + "/api/data/text")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.TextType)
	suite.Require().Equal(result.Name, "text")
	suite.Require().Equal(result.Data, "some data")
}

func (suite *ServerTestSuite) TestAddBinaryData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		AddData(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "binary",
				Data:      "url",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	mockMinio.EXPECT().
		CreateOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			"123", nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetFile("file", "fixtures/test_image.webp").
		SetFormData(map[string]string{
			"name": "binary",
		}).
		SetResult(&result).
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Authorization", "Bearer "+suite.token).
		Post(ts.URL + "/api/data/binary")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.BinaryType)
}

func (suite *ServerTestSuite) TestUpdateTextData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)
	mockDB.EXPECT().
		UpdateData(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse
	body := model.AddTextData{
		Name: "text",
		Data: "some data",
	}

	resp, err := suite.client.R().
		SetBody(body).
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Put(ts.URL + "/api/data/text/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.TextType)
	suite.Require().Equal(result.Name, "text")
	suite.Require().Equal(result.Data, "some data")
}

func (suite *ServerTestSuite) TestUpdateBinaryData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "binary",
				Data:      "url",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)
	mockDB.EXPECT().
		UpdateData(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "binary",
				Data:      "some data",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	mockMinio.EXPECT().
		DeleteOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)
	mockMinio.EXPECT().
		GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			"123", nil)
	mockMinio.EXPECT().
		CreateOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			"123", nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetFile("file", "fixtures/test_image.webp").
		SetFormData(map[string]string{
			"name": "binary",
		}).
		SetResult(&result).
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("Authorization", "Bearer "+suite.token).
		Put(ts.URL + "/api/data/binary/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.BinaryType)
}

func (suite *ServerTestSuite) TestDeleteTextData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)
	mockDB.EXPECT().
		DeleteData(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "text",
				Data:      "some data",
				Type:      model.TextType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Delete(ts.URL + "/api/data/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.TextType)
	suite.Require().Equal(result.Name, "text")
	suite.Require().Equal(result.Data, "some data")
}

func (suite *ServerTestSuite) TestDeleteBinaryData() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := mock.NewMockKeeperStorage(ctrl)
	mockMinio := mock.NewMockMinioStorage(ctrl)
	mockDB.EXPECT().
		GetDataByID(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "binary",
				Data:      "url",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)
	mockDB.EXPECT().
		DeleteData(gomock.Any(), gomock.Any()).
		Return(
			model.Keeper{
				ID:        1,
				Name:      "binary",
				Data:      "some data",
				Type:      model.BinaryType,
				UserID:    userID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	mockMinio.EXPECT().
		DeleteOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)
	mockMinio.EXPECT().
		GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			"123", nil)

	keeperService := keeper.NewService(
		keeper.WithConfig(suite.config),
		keeper.WithStorage(mockDB),
		keeper.WithMinio(mockMinio),
	)
	handlers := handler.NewHandler(
		handler.WithConfig(suite.config),
		handler.WithValidator(sharedUtils.NewValidator()),
		handler.WithKeeperService(keeperService),
	)

	ts := httptest.NewServer(router.NewRouter(handlers))
	defer ts.Close()

	var result model.DataResponse

	resp, err := suite.client.R().
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+suite.token).
		Delete(ts.URL + "/api/data/1")
	suite.Require().NoError(err)
	suite.Require().Equal(200, resp.StatusCode())
	suite.Require().Equal(result.Type, model.BinaryType)
}
