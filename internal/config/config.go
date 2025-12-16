package config

import "os"

// Config holds application configuration values
type Config struct {
	HTTPAddr     string
	LogLevel     string
	MQTTBroker   string
	MQTTClientID string
}

// Load reads configuration from environment variables
func Load() Config {
	return Config{
		HTTPAddr:     getenv("HTTP_ADDR", ":8080"),
		LogLevel:     getenv("LOG_LEVEL", "info"),
		MQTTBroker:   getenv("MQTT_BROKER", "tcp://localhost:1883"),
		MQTTClientID: getenv("MQTT_CLIENT_ID", "orion-dev"),
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
