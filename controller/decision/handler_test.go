package decision

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/labstack/echo/v4"
	// "github.com/sanglx-teko/opa-dispatcher/config"
	// "github.com/sanglx-teko/opa-dispatcher/cores/configurationmanager"
	"github.com/sanglx-teko/opa-dispatcher/model"
	"github.com/stretchr/testify/assert"
)

func GetServiceConfigMock() (result map[string]*model.ServiceGroup) {
	config := `{ 
		"iam": {
		  "name": "core services",
		  "url": "http://localhost:8181/v1/data/rbac/authz/allow"
		},
		"online": {
		  "name": "online services",
		  "url": "http://localhost:8182/v1/data/rbac/authz/allow"
		}
	  }`
	if err := json.Unmarshal([]byte(config), &result); err != nil {
		result = nil
	}
	return
}

func HandleDecisionAPIControllerMock(c echo.Context) (erro error) {
	response := &ResponseData{}
	defer func() {
		if erro == nil {
			erro = c.JSON(http.StatusOK, response)
			return
		}
	}()

	u := new(model.DecisionRequest)
	if err := c.Bind(u); err != nil {
		response.SetAllowed(false)
		return
	}

	serviceConfig := GetServiceConfigMock()
	val, ok := serviceConfig[u.Service]

	if !ok {
		response.SetAllowed(false)
		return
	}
	type tempStruct struct {
		Input *model.DecisionRequest `json:"input"`
	}

	requestBody, _ := json.Marshal(&tempStruct{
		Input: u,
	})

	resp, err := http.Post(val.URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		response.SetAllowed(false)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.SetAllowed(false)
		return
	}
	type OPAResp struct {
		Result bool `json:"result"`
	}
	var temp *OPAResp
	if err := json.Unmarshal(body, &temp); err != nil {
		response.SetAllowed(false)
		return
	}
	response.SetAllowed(temp.Result)
	return
}

func TestHandleDecisionAPIController(t *testing.T) {
	e := echo.New()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://localhost:8181/v1/data/rbac/authz/allow",
		httpmock.NewStringResponder(200, `{"result": true}`))
	httpmock.RegisterResponder("POST", "http://localhost:8182/v1/data/rbac/authz/allow",
		httpmock.NewStringResponder(200, `{"result": false}`))

	temp := model.DecisionRequest{
		Action:   "write",
		Service:  "online",
		Subject:  "U01",
		Resource: "permissions",
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
	if assert.NoError(t, HandleDecisionAPIControllerMock(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		response := &ResponseData{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		}
		assert.Equal(t, false, response.Allowed)
	}

}
