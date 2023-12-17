package app

import "product_api/common/postgresql"

type ConfigurationManager struct {
	PostgresSqlConfig postgresql.Config
}

func NewConfiguraitonManager() *ConfigurationManager {
	pSqlConfig := getPSqlConfig()
	return &ConfigurationManager{PostgresSqlConfig: pSqlConfig}
}
func getPSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productdb",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "100s",
	}
}
