package handlers_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"errors"

	"product/handlers"
	"product/models"
	"product/utils/errs"
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

func TestGetIntId(t *testing.T) {
	e := echo.New()

	t.Run("test case : pass valid integer id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("123")

		id, err := handlers.GetIntId(c)

		assert.NoError(t, err)
		assert.Equal(t, 123, id)
	})

	t.Run("test case : invalid non-integer id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("a")

		id, err := handlers.GetIntId(c)

		expectedErr := errs.NewBadRequestError("Invalid id: a is not integer")

		assert.Error(t, err)
        assert.Equal(t, expectedErr, err)
        assert.Equal(t, 0, id)
	})
}

func TestProductValidator_Validate(t *testing.T) {
	validator := handlers.NewProductValidator()

	t.Run("test case : valid input", func(t *testing.T) {
		productCreate := models.ProductCreate{
			Id: 1,
			ProdType_Id: 1,
			Name: "burger",
			Desc: "hey",
			Price: 100.0,
			Discount: 0.0,
			CreatedUser: 1,
		}

		err := validator.Validate(productCreate)

		assert.NoError(t, err)
	})

	t.Run("test case : invalid input", func(t *testing.T) {
		invalidProductCreate := models.ProductCreate{
			Id: 1,
			ProdType_Id: 1,
			Name: "",
			Desc: "hey",
			Price: 0,
			Discount: 0.0,
			CreatedUser: 1,
		}

		err := validator.Validate(invalidProductCreate)

		assert.Error(t, err)
	})
}

func TestProductTypeValidator_Validate(t *testing.T) {
	validator := handlers.NewProductTypeValidator()

	t.Run("valid input", func(t *testing.T) {
		productTypeCreate := models.ProductTypeCreate{
			Id: 1,
			Name: "burger",
		}

		err := validator.Validate(productTypeCreate)

		assert.NoError(t, err)
	})

	t.Run("invalid input", func(t *testing.T) {
		invalidProductTypeCreate := models.ProductTypeCreate{
			Id: 1,
			Name: "",
		}

		err := validator.Validate(invalidProductTypeCreate)

		assert.Error(t, err)
	})
}