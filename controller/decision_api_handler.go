package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DecisionRequest struct {
	Action   string `json:"action" form:"action" query:"action"`
	Resource string `json:"resource" form:"resource" query:"resource"`
	Subject  string `json:"subject" form:"subject" query:"subject"`
	Service  string `json:"service" form:"service" query:"service"`
}

// HandleDecisionAPIController ...
func HandleDecisionAPIController(c echo.Context) error {
	u := new(DecisionRequest)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, u)
}
