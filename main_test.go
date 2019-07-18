package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		res := map[string]string{}
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {

		}
		assert.Equal(t, map[string]string{"Hello": "World"}, res)
	}
}
