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
	assert.Equal(t, "thomas", AppConfig.Database.Dsn)
	assert.Equal(t, 11, AppConfig.Database.MaxIdleConns)
	assert.Equal(t, 114, AppConfig.Database.MaxOpenConns)
}
