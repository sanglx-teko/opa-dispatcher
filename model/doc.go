package model

// DecisionRequest ...
type DecisionRequest struct {
	Action   string `json:"action" form:"action" query:"action"`
	Resource string `json:"resource" form:"resource" query:"resource"`
	Subject  string `json:"subject" form:"subject" query:"subject"`
	Service  string `json:"service" form:"service" query:"service"`
}

// ServiceGroup ...
type ServiceGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
