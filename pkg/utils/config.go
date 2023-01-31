package utils

func GetConfigPath(configPath string) string {
	if configPath == "docker" || configPath == "debug" {
		return "config/config-local.yml"
	}
	return "./config/config-local"
}
