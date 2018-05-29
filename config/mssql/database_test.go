package mssql

import (
	"go/build"
	"testing"

	config "github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMsSQLDBConnection(t *testing.T) {
	envDir := build.Default.GOPATH + "/src/github.com/Bhinneka/b2c-api/"
	err := config.Load(envDir + ".env")
	if err != nil {
		assert.Error(t, err)
	}

	if testing.Short() {
		t.Skip("Skipping Integration Test on Short Mode")
	}

	t.Run("TestAuthMsSQLDBConnection", func(t *testing.T) {
		db := GetAuthMsSQLDB()

		err := db.Ping()

		assert.NoError(t, err)
	})

	t.Run("TestLegacyMsSQLDBConnection", func(t *testing.T) {
		db := GetLegacyMsSQLDB()

		err := db.Ping()

		assert.NoError(t, err)
	})
}
