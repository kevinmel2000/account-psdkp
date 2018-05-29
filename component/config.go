package component

import (
	"fmt"
	"log/syslog"
	"os"
	"strings"
	log "github.com/sirupsen/logrus"
	logrusSyslog "github.com/sirupsen/logrus/hooks/syslog"
	config "github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
	"github.com/Bhinneka/api-gateway-apps/model"
)

// InitConfig function for initializing config data
// configDir string directory where the config is located
func InitConfig(configDir string) error {
	if strings.TrimSpace(configDir) == "" {
		message := "config_dir is not defined"
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow(message))
		os.Exit(1)
	}
	configDir = strings.TrimRight(configDir, "/")

	err := config.Load()
	if err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow(err.Error()))
	}

	os.Setenv("DIR", configDir)

	return nil
}

// InitLogger function for initializing log data
func InitLogger() error {
	log.SetFormatter(&log.JSONFormatter{})
	syslogOutput, err := logrusSyslog.NewSyslogHook("", "", syslog.LOG_INFO, model.LogTag)
	log.AddHook(syslogOutput)

	if err != nil {
		log.Fatal("Unable to setup syslog output")
	}

	logLevel := log.WarnLevel
	if os.Getenv("DEBUG") == "1" {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	return nil
}


// Init function for initializing config data logger
// configDir string directory where the config is located
func Init(configDir string) error {
	if err := InitConfig(configDir); err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow(err.Error()))
	}

	// call another initialization here
	if err := InitLogger(); err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow(err.Error()))
	}
	//InitDB()

	return nil
}
