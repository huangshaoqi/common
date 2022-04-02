package common

import (
	"github.com/asim/go-micro/v3/config"
	log "github.com/asim/go-micro/v3/logger"
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
	err := config.Get(path).Scan(mysqlConfig)
	if err != nil {
		log.Error(err)
		return nil
	}
	return mysqlConfig
}
