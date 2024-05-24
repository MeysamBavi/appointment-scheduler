package app_test

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

const testDBName = "test_gorm.db"

type IntegrationTestSuite struct {
	suite.Suite

	service *app.HTTPService

	jwt    *jwt.JWT
	gormDB *gorm.DB
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (t *IntegrationTestSuite) SetupTest() {
	gormDB, err := gorm.Open(sqlite.Open(testDBName), &gorm.Config{TranslateError: true})
	t.NoError(err)
	t.gormDB = gormDB

	jwtSdk := jwt.NewJWT("sth")
	t.jwt = jwtSdk

	t.service = app.NewHTTPService(app.Config{Port: 8888}, jwtSdk, gormDB)

	result := t.gormDB.Create(&models.Business{
		Model:   gorm.Model{ID: 1},
		Name:    "ali",
		Address: "tehran",
		UserID:  10,
	})
	t.NoError(result.Error)
}

func (t *IntegrationTestSuite) TearDownTest() {
	err := os.Remove(testDBName)
	t.NoError(err)
}

func (t *IntegrationTestSuite) TestCreateEmployeeCreatesEmployeeWithoutError() {
	body := `{"user": 5}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/1/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusCreated, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("employee created.", bodyMap["message"].(string))
}

func (t *IntegrationTestSuite) TestCreateEmployeeReturnErrorWhenBusinessNotExists() {
	body := `{"user": 5}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/2/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("2")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusNotFound, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("business not found.", bodyMap["message"].(string))
}

func (t *IntegrationTestSuite) TestCreateEmployeeReturnErrorRequesterIsNotOwner() {
	body := `{"user": 5}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/1/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(222))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusForbidden, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("you aren't business owner.", bodyMap["message"].(string))
}

func (t *IntegrationTestSuite) TestCreateEmployeeReturnErrorWhenRequestIsDuplicate() {
	body := `{"user": 50}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/2/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	req = httptest.NewRequest(http.MethodPost, "/businesses/2/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)
	ctx.Set("userId", int32(10))
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err = t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusConflict, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("employee already exist.", bodyMap["message"].(string))
}

func (t *IntegrationTestSuite) TestCreateEmployeeReturnErrorWhenNotAuthorized() {
	body := `{"user": 50}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/2/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusUnauthorized, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("you are not authorized.", bodyMap["message"].(string))
}

func (t *IntegrationTestSuite) TestCreateEmployeeReturnErrorWhenUserNotSent() {
	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/businesses/2/employees", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("business_id")
	ctx.SetParamValues("1")
	err := t.service.CreateEmployee(ctx)
	t.NoError(err)

	t.Equal(http.StatusBadRequest, rec.Result().StatusCode)

	bodyMap := make(map[string]any)
	err = json.Unmarshal(rec.Body.Bytes(), &bodyMap)
	t.NoError(err)
	t.Equal("you should send user id to create employee.", bodyMap["message"].(string))
}
