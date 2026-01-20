package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladislavkovaliov/ledger/internal/config"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "1234")
	os.Setenv("DATABASE_URL", "localhost:27017")

	cfg := config.LoadConfig()

	assert.Equal(t, "1234", cfg.Port)
	assert.Equal(t, "localhost:27017", cfg.DatabaseUrl)
}
