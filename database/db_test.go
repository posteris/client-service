package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/database"
)

func TestInitDatabase(t *testing.T) {
	//obtains the connection database parameters from the ci-database-starter
	//test module
	tests := conndata.GetTestData()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			os.Setenv(database.DatabaseTypeLabel, tt.Type)
			os.Setenv(database.DatabaseDsnLabel, tt.Args.DSN)

			InitDatabase()

			assert.NotNil(t, GetInstance())
		})
	}
}
