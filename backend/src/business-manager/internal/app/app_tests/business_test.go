package app_tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/jwt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/app"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const businessTestDBName = "test_gorm.db"

type BusinessTestSuite struct {
	suite.Suite

	service *app.HTTPService

	jwt    *jwt.JWT
	gormDB *gorm.DB
}

func TestBusinessTestSuite(t *testing.T) {
	suite.Run(t, new(BusinessTestSuite))
}

func (t *BusinessTestSuite) SetupTest() {
	gormDB, err := gorm.Open(sqlite.Open(businessTestDBName), &gorm.Config{TranslateError: true})
	t.NoError(err)
	t.gormDB = gormDB

	jwtSdk := jwt.NewJWT("sth")
	t.jwt = jwtSdk

	t.service = app.NewHTTPService(app.Config{Port: 8888}, jwtSdk, gormDB)

	result := t.gormDB.Create(&models.ServiceType{
		Model: gorm.Model{ID: 50},
		Name:  "st1",
	})
	t.NoError(result.Error)
}

func (t *BusinessTestSuite) TearDownTest() {
	err := os.Remove(businessTestDBName)
	t.NoError(err)
}

func (t *BusinessTestSuite) TestCreateBusinessCreatesBusinessWithoutError() {
	body := `{"name": "b1", "address": "a1", "service_type": 50}`
	req := httptest.NewRequest(http.MethodPost, "/businesses", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateBusiness(ctx)
	t.NoError(err)

	t.Equal(http.StatusCreated, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("business created.", bodyMap["message"].(string))
}

func (t *BusinessTestSuite) TestCreateBusinessReturnErrorIfServiceTypeIsInvalid() {
	body := `{"name": "b1", "address": "a1", "service_type": 500}`
	req := httptest.NewRequest(http.MethodPost, "/businesses", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateBusiness(ctx)
	t.NoError(err)

	t.Equal(http.StatusNotFound, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("service type not found.", bodyMap["message"].(string))
}
