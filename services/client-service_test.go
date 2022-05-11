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

//obtains the connection database parameters from the ci-database-starter
var databases = conndata.GetTestData()

func TestCreate(t *testing.T) {

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
					name: "success",
					args: args{
						client: &models.Client{
							Name:    uuid.NewString(),
							Surname: uuid.NewString(),
							Email:   fmt.Sprintf("%s@test.go", uuid.NewString()),
						},
					},
					wantErr: false,
				},
				{
					name: "fail",
					args: args{
						client: &models.Client{
							Name:  uuid.NewString(),
							Email: fmt.Sprintf("%s@test.go", uuid.NewString()),
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

func TestGetById(t *testing.T) {
	for _, dbData := range databases {
		os.Setenv(database.DatabaseTypeLabel, dbData.Type)
		os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

		db.InitDatabase()
		db.Automigrate()

		client := models.Client{
			Name:    uuid.NewString(),
			Surname: uuid.NewString(),
			Email:   fmt.Sprintf("%s@test.go", uuid.NewString()),
		}

		err := Create(&client)

		if err != nil {
			t.Error("Unable to create client")
		}

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
					id: client.ID,
				},
				want: &client,
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
			t.Run(fmt.Sprintf("%s-%s", dbData.Name, tt.name), func(t *testing.T) {
				got := GetById(tt.args.id)

				if got == nil && tt.want == nil {
					return
				}

				//due to clickhouse fixedstring strategy, I'd cretae it to clean and
				//compare it
				compareString := func(gotString string, wantString string) bool {
					if dbData.Type != database.Clickhouse {
						return gotString == wantString
					}

					wantByte := []byte(wantString)
					gotByte := []byte(gotString)

					size := len(wantByte)

					gotByte = gotByte[:size]

					return string(gotByte) == wantString
				}

				idEquals := got.ID == tt.want.ID
				nameEquals := compareString(got.Name, tt.want.Name)
				surnameEquals := compareString(got.Surname, tt.want.Surname)
				emailEquals := compareString(got.Email, tt.want.Email)

				equals := idEquals && nameEquals && surnameEquals && emailEquals

				if !equals {
					t.Errorf("GetById() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
