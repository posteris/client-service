package services

import (
	"os"
	"testing"

	"github.com/google/uuid"
	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/client-service/db"
	"github.com/posteris/client-service/models"
	"github.com/posteris/database"
)

func TestGetById(t *testing.T) {
	//obtains the connection database parameters from the posteris/ci-database-starter
	var databases = conndata.GetTestData()

	//execute the all test cases for each database
	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			//set environment variables
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			//start database
			db.InitDatabase()

			cli := createNewClientToUpdate()

			type args struct {
				id string
			}
			tests := []struct {
				name string
				args args
				want *models.Client
			}{
				{
					name: "success",
					args: args{
						id: cli.ID,
					},
					want: cli,
				},
				{
					name: "error",
					args: args{
						id: uuid.NewString(),
					},
					want: nil,
				},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					found := GetById(tt.args.id)

					errorNilNotFound := found != nil && tt.want == nil
					errorElementNotFound := found == nil && tt.want != nil

					if errorNilNotFound || errorElementNotFound {
						t.Errorf("GetById() = %v, want %v", found, tt.want)
					}
				})
			}
		})
	}
}
