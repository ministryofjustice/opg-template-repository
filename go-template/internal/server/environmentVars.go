package server

import "os"

type EnvironmentVars struct {
	Port            string
	WebDir          string
	Prefix          string
	SiriusPublicURL string
}

func NewEnvironmentVars() EnvironmentVars {
	return EnvironmentVars{
		Port:            getEnv("PORT", "{{PORT}}"),
		WebDir:          getEnv("WEB_DIR", "web"),
		SiriusPublicURL: getEnv("SIRIUS_PUBLIC_URL", ""),
		Prefix:          getEnv("PREFIX", "{{URL_PREFIX}}"),
	}
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
