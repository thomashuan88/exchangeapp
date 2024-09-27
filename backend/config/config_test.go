package config

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInitConfig(t *testing.T) {
	err := os.Chdir("..")
	if err != nil {
		t.Fatal(err)
	}
	InitConfig()
	assert.Equal(t, "CurrencyExchangeApp", AppConfig.App.Name)
	assert.Equal(t, "8000", AppConfig.App.Port)
	assert.Equal(t, "localhost", AppConfig.Database.Host)
	assert.Equal(t, "3306", AppConfig.Database.Port)
	assert.Equal(t, "thomas", AppConfig.Database.User)
	assert.Equal(t, "huan1122", AppConfig.Database.Password)
	assert.Equal(t, "currency_exchange", AppConfig.Database.Dbname)
}
