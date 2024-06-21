package handlers_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"errors"
	"time"

	"timestamp/handlers"
	"timestamp/utils/errs"
)

func TestHandleError(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		appErr := errs.AppError{Code: http.StatusBadRequest, Message: "Bad Request"}

		handlers.HandleError(e.NewContext(req, rec), appErr)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		expectedBody := `{"code":400,"message":"Bad Request"}`
		actualBody := strings.TrimSpace(rec.Body.String())

		assert.Equal(t, expectedBody, actualBody)
	})

	t.Run("test case : fail", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		appErr := errors.New("")

		handlers.HandleError(e.NewContext(req, rec), appErr)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expectedBody := `{"code":500,"message":"Interval Server Error"}`
		actualBody := strings.TrimSpace(rec.Body.String())

		assert.Equal(t, expectedBody, actualBody)
	})
}

func TestGetParamDay(t *testing.T) {
	e := echo.New()

	t.Run("test case : pass valid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2016-12-01")

		date, err := handlers.GetParamDay(c)

		expected := time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC)

		assert.NoError(t, err)
		assert.Equal(t, expected, date)
	})

	t.Run("test case : invalid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("a")

		date, err := handlers.GetParamDay(c)

		expectedErr := errs.NewBadRequestError("Invalid date format")

		assert.Error(t, err)
        assert.Equal(t, expectedErr, err)
        assert.Equal(t, time.Time{}, date)
	})
}

func TestGetParamMonth(t *testing.T) {
	e := echo.New()

	t.Run("test case : pass valid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2016-12")

		date, err := handlers.GetParamMonth(c)

		expected := time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC)

		assert.NoError(t, err)
		assert.Equal(t, expected, date)
	})

	t.Run("test case : invalid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("a")

		date, err := handlers.GetParamMonth(c)

		expectedErr := errs.NewBadRequestError("Invalid date format")

		assert.Error(t, err)
        assert.Equal(t, expectedErr, err)
        assert.Equal(t, time.Time{}, date)
	})
}

func TestGetParamYear(t *testing.T) {
	e := echo.New()

	t.Run("test case : pass valid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2016")

		date, err := handlers.GetParamYear(c)

		expected := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)

		assert.NoError(t, err)
		assert.Equal(t, expected, date)
	})

	t.Run("test case : invalid date format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("a")

		date, err := handlers.GetParamYear(c)

		expectedErr := errs.NewBadRequestError("Invalid date format")

		assert.Error(t, err)
        assert.Equal(t, expectedErr, err)
        assert.Equal(t, time.Time{}, date)
	})
}