package decision

import "github.com/sanglx-teko/opa-dispatcher/cores/configurationmanager/model"

// ConfigurationManager ...
var ConfigurationManager model.IConfigurationManager

// InitCFManager ...
func InitCFManager(conf model.IConfigurationManager) {
	ConfigurationManager = conf
}

type (
	// ResponseData ...
	ResponseData struct {
		Allowed bool `json:"allowed"`
	}
)

// SetAllowed ...
func (resp *ResponseData) SetAllowed(allowed bool) {
	resp.Allowed = allowed
}
