package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Init(t *testing.T) {
	Init()
	actual := viper.GetString(connectionStringName)
	expected := defaultConnectString
	assert.Equal(t, expected, actual, "Must be equal default value")
}

func Test_ReadConnectionString(t *testing.T) {
	expected := "100500"
	err := os.Setenv(connectionStringName, expected)
	if err != nil {
		t.Fatal(err)
	}
	Init()
	actual := ReadConnectionString()
	assert.Equal(t, expected, actual, "Must be read value from environment")
}
