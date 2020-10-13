package sqlite

import (
	log "github.com/sirupsen/logrus"
)

// Create plugin with Logrus
type Logger struct{}

// Print implements log level
func (c *Logger) Print(v ...interface{}) {
	switch v[0] {
	case "error":
		log.WithFields(log.Fields{"provider": "sqlite", "file": v[1]}).Error(v[2])
	case "sql":
		log.WithFields(
			log.Fields{
				"provider": "sqlite",
				"values":   v[4],
				"returned": v[5],
				"duration": v[2],
			},
		).Trace(v[3])
	case "log":
		log.WithFields(log.Fields{"provider": "sqlite", "file": v[1]}).Warn(v[2])
	}
}
