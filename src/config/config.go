package config

import (
	"github.com/spf13/viper"
)

const connectionStringName = "DATABASE_CONNECTION"
const defaultConnectString = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func Init() {
	viper.SetDefault(connectionStringName, defaultConnectString)
	err := viper.BindEnv(connectionStringName)
	if err != nil {
		panic(error.Error)
	}
}

func ReadConnectionString() string {
	return viper.GetString(connectionStringName)
}
