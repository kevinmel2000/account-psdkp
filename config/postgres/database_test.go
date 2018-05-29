package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresDBConnection(t *testing.T) {
	//envDir := build.Default.GOPATH + "/src/github.com/Bhinneka/b2c-api/"
	//err := config.Load(envDir + ".env")
	//if err != nil {
	//	assert.Error(t, err)
	//}

	if testing.Short() {
		t.Skip("Skipping Integration Test on Short Mode")
	}

	t.Run("TestWritePostgresDBConnection", func(t *testing.T) {
		db := WritePostgresDB()

		err := db.Ping()

		assert.NoError(t, err)

	})

	t.Run("TestReadPostgresDBConnection", func(t *testing.T) {
		db := ReadPostgresDB()

		err := db.Ping()

		assert.NoError(t, err)

	})
}
