package services

import (
	"github.com/dezzles-apps/go-common/db"
)

type ConfigService struct {
	connection *db.Database
}

func NewConfigService(connection *db.Database) *ConfigService {
	return &ConfigService{
		connection: connection,
	}
}

func (cs *ConfigService) getConfigValue(key string) (string, error) {
	var value string
	err := cs.connection.GetDB().QueryRow("SELECT value FROM config WHERE config_key = ?", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (cs *ConfigService) CanRegisterUsers() bool {
	value, err := cs.getConfigValue("register_enabled")
	if err != nil {
		return false
	}
	return value == "true"
}
