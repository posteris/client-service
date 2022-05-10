package db

import (
	"os"
	"os/exec"
	"testing"

	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/database"
	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	//obtains the connection database parameters from the ci-database-starter
	databases := conndata.GetTestData()

	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			InitDatabase()

			assert.NotNil(t, GetInstance())
		})
	}
}

func TestInitDatabaseCrash(t *testing.T) {
	if os.Getenv("CRASH") == "1" {
		os.Setenv(database.DatabaseTypeLabel, "error")
		os.Setenv(database.DatabaseDsnLabel, "error")

		InitDatabase()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitDatabaseCrash")
	cmd.Env = append(os.Environ(), "CRASH=1")
	err := cmd.Run()

	e, ok := err.(*exec.ExitError)

	assert.True(t, ok)
	assert.False(t, e.Success())
}

func TestAutomigrate(t *testing.T) {
	//obtains the connection database parameters from the ci-database-starter
	databases := conndata.GetTestData()

	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			InitDatabase()
			Automigrate()

			assert.NotNil(t, GetInstance())
		})
	}
}
