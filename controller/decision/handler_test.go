package decision

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sanglx-teko/opa-dispatcher/model"
	"github.com/stretchr/testify/assert"
)

func TestHandleDecisionAPIController(t *testing.T) {
	e := echo.New()
	temp := model.DecisionRequest{
		Action:   "write",
		Service:  "online",
		Subject:  "U03",
		Resource: "wallets",
	}
	res := []byte{}
	res, err := json.Marshal(temp)
	if err != nil {
		t.Error(err)
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/decision/handler", bytes.NewReader(res))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, HandleDecisionAPIController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
