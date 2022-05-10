package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/client-service/db"
	"github.com/posteris/client-service/models"
	"github.com/posteris/database"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	//obtains the connection database parameters from the ci-database-starter
	databases := conndata.GetTestData()

	for _, dbData := range databases {
		os.Setenv(database.DatabaseTypeLabel, dbData.Type)
		os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

		db.InitDatabase()
		db.Automigrate()

		t.Run(dbData.Name, func(t *testing.T) {
			type args struct {
				client *models.Client
			}
			tests := []struct {
				name    string
				args    args
				wantErr bool
			}{
				{
					name: "create",
					args: args{
						client: &models.Client{
							Name:    uuid.NewString(),
							Surname: uuid.NewString(),
							Email:   fmt.Sprintf("%s@test.go", uuid.NewString()),
						},
					},
					wantErr: false,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if err := Create(tt.args.client); (err != nil) != tt.wantErr {
						t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					}

					instance := db.GetInstance()

					var client models.Client
					instance.First(&client, "email = ?", tt.args.client.Email)

					assert.NotNil(t, client)
				})
			}
		})
	}
}
