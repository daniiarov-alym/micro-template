package config

import (
	"os"
	"strings"

	logger "github.com/sirupsen/logrus"
)

var config Cfg

type Cfg struct {
	DbHost		string
	DbPort		string
	DbUser		string
	DbPassword	string
	DbName 		string
	ApiPort		string
}

func ReadConfig()  {
	missing := make([]string, 0)
	config.DbHost = ReadField("PG_HOST", &missing)
	config.DbPort = ReadField("PG_PORT", &missing)
	config.DbUser = ReadField("PG_USER", &missing)
	config.DbPassword = ReadField("PG_PASSWORD", &missing)
	config.DbName = ReadField("PG_DATABASE", &missing)
	config.ApiPort = ReadField("PORT", &missing)
	if len(missing) != 0 {
		logger.Fatal(strings.Join(missing, ", ") + " fields are missing")
	}
}

func ReadField(key string, absent *[]string) string {
	val := os.Getenv(key)
	if (val == "") {
		*absent = append(*absent, key)
	}
	return val
}

func Config() *Cfg {
	return &config
}