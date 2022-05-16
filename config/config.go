package config

import (
	"os"
	"strconv"
)

// MailerConfig is a structure for SMTP environment variables
type MailerConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

// InitConfig is a function for initialize Config from environment variables
func InitConfig() *MailerConfig {
	return &MailerConfig{
		SMTPHost:     getEnv("SMTP_HOST", ""),
		SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
		SMTPUsername: getEnv("SMTP_USERNAME", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
