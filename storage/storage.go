package storage

import (
	"encoding/json"
	"github.com/uadmin/uadmin"
	"os"
)

type config struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	DbName   string `json:"DbName"`
	Type     string `json:"Type"`
}

func InitDB(pathToConf string) (*uadmin.DBSettings, error) {
	configFile, err := os.Open(pathToConf)
	if err != nil {
		return nil, err
	}

	dbConfig := &config{}
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(dbConfig); err != nil {
		return nil, err
	}

	return &uadmin.DBSettings{
		Type:     dbConfig.Type,
		Name:     dbConfig.DbName,
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		User:     dbConfig.User,
		Password: dbConfig.Password,
	}, nil
}
