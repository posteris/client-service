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
)

func TestCreate(t *testing.T) {
	//obtains the connection database parameters from the posteris/ci-database-starter
	var databases = conndata.GetTestData()

	//create test cases
	type args struct {
		client models.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "lower-case",
			args: args{
				client: models.Client{
					Name:    models.GetStringPointer("john"),
					Surname: models.GetStringPointer("smith"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@test.go", uuid.NewString()),
					),
				},
			},
			wantErr: false,
		},
		{
			name: "upper-case",
			args: args{
				client: models.Client{
					Name:    models.GetStringPointer("JOHN"),
					Surname: models.GetStringPointer("SMITH"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@TEST.GO", uuid.NewString()),
					),
				},
			},
			wantErr: false,
		},
		{
			name: "no-name",
			args: args{
				client: models.Client{
					Surname: models.GetStringPointer("SMITH"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@TEST.GO", uuid.NewString()),
					),
				},
			},
			wantErr: true,
		},
		{
			name: "no-surname",
			args: args{
				client: models.Client{
					Name: models.GetStringPointer("JOHN"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@TEST.GO", uuid.NewString()),
					),
				},
			},
			wantErr: true,
		},
		{
			name: "no-email",
			args: args{
				client: models.Client{
					Name:    models.GetStringPointer("JOHN"),
					Surname: models.GetStringPointer("SMITH"),
				},
			},
			wantErr: true,
		},
		{
			name: "activated",
			args: args{
				client: models.Client{
					Name:    models.GetStringPointer("JOHN"),
					Surname: models.GetStringPointer("SMITH"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@TEST.GO", uuid.NewString()),
					),
					Active: true,
				},
			},
			wantErr: true,
		},
		{
			name: "with-id",
			args: args{
				client: models.Client{
					ID:      uuid.NewString(),
					Name:    models.GetStringPointer("JOHN"),
					Surname: models.GetStringPointer("SMITH"),
					Email: models.GetStringPointer(
						fmt.Sprintf("%s@TEST.GO", uuid.NewString()),
					),
				},
			},
			wantErr: true,
		},
	}

	//execute the all test cases for each database
	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			//set environment variables
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			//start database
			db.InitDatabase()

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					err := Create(&tt.args.client)

					if (err != nil) != tt.wantErr {
						t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					}

					//check if object was created
					if err == nil {
						client := GetById(tt.args.client.ID)

						if client == nil {
							t.Errorf("Unable to create the Client %s", client.ToString())
						}
					}
				})
			}
		})
	}
}
