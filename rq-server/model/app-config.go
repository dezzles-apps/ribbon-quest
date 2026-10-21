package model

type AppConfig struct {
	JwtSecret string `yaml:"jwt_secret"`
	DebugLog  bool   `yaml:"debugLogging"`
}
