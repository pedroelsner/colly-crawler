package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

// Define some basics configs for application:
//  - Load environment vars
//  - Custom logs format/level
func config() {
	// Load .env file
	_ = godotenv.Load()

	// Output to stdout instead of the default stderr
	logrus.SetOutput(os.Stdout)

	// Set Defaults
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys: true,
		FieldsOrder: []string{
			"provider",
			"url",
			"duration",
		},
	})

	// When production:
	//  - Log as JSON instead of the default ASCII formatter
	//  - Set level error to Info
	if os.Getenv("APP_ENV") == "production" {
		logrus.SetLevel(logrus.InfoLevel)

		// Change field "msg" to "message"
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		})
	}
}
