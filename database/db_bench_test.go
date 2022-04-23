package database

import (
	"os"
	"testing"

	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/database"
)

func BenchmarkInitDatabase(b *testing.B) {

	//obtains the connection database parameters from the ci-database-starter
	//test module
	tests := conndata.GetTestData()

	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			os.Setenv(database.DatabaseTypeLabel, tt.Type)
			os.Setenv(database.DatabaseDsnLabel, tt.Args.DSN)

			InitDatabase()
		})
	}
}
