package decision

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sanglx-teko/opa-dispatcher/model"
)

// HandleDecisionAPIController ...
func HandleDecisionAPIController(c echo.Context) (erro error) {
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

	// service: "iam"
	serviceConfig := ConfigurationManager.GetServiceConfig()
	val, ok := serviceConfig[u.Service]

	if !ok {
		response.SetAllowed(false)
		return
	}

	type tempStruct struct {
		Input *model.DecisionRequest `json:"input"`
	}

	requestBody,_ := json.Marshal(&tempStruct{
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
