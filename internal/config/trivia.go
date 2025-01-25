package config

// Config retrieves a configuration value by key
func Config(key string) string {
	// ConfigMap stores configuration values
	var ConfigMap = map[string]string{
		"APP_ENV":   GetEnvWithKey("APP_ENV", "production"),
		"APP_DEBUG": GetEnvWithKey("APP_DEBUG", "false"),
		"APP_NAME":  GetEnvWithKey("APP_NAME", "Go App"),
		"APP_URL":   GetEnvWithKey("APP_URL", "http://localhost"),
		"PORT":      GetEnvWithKey("PORT", "8000"),
	}
	return ConfigMap[key]
}
