package main

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sanglx-teko/opa-dispatcher/config"
	"github.com/sanglx-teko/opa-dispatcher/controller/decision"
	manager "github.com/sanglx-teko/opa-dispatcher/cores/configurationmanager"
)

func changeWorkingDir() (currentDir string, err error) {
	currentDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	err = os.Chdir(currentDir)
	return
}

// loadConfigurationAndDB load configuration and setup Redis, DB, RPC, etc
func loadConfigurationAndDB(currentDir string) error {
	// Load configurations
	if err := config.LoadConfigurations("configs.json"); err != nil {
		return err
	}

	return nil
}

func initRouter() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Remove trailing slash middleware removes a trailing slash from the request URI.
	e.Pre(middleware.RemoveTrailingSlash())

	// Secure middleware provides protection against cross-site scripting (XSS) attack, content type sniffing, clickjacking, insecure connection and other code injection attacks.
	// For more example, please refer to https://echo.labstack.com/
	e.Use(middleware.Secure())

	e.POST("/decision/handler", decision.HandleDecisionAPIController)
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	currentDir, err := changeWorkingDir()
	if err != nil {
		panic(err)
	}

	if err := loadConfigurationAndDB(currentDir); err != nil {
		panic(err)
	}

	manager.Instance.InitWithConfig(config.GetConfigurations().ETCD)
	decision.InitCFManager(manager.Instance)

	initRouter()
}
