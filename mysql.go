package common

import (
	"github.com/micro/micro/v3/service/config"
	log "github.com/micro/micro/v3/service/logger"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	val, err := config.Get(path)
	if err != nil {
		log.Error(err)
		return nil
	}

	err = val.Scan(mysqlConfig)
	if err != nil {
		log.Error(err)
		return nil
	}
	return mysqlConfig
}
