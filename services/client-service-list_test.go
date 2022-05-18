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

func TestList(t *testing.T) {
	//obtains the connection database parameters from the posteris/ci-database-starter
	var databases = conndata.GetTestData()

	type args struct {
		search       map[string]interface{}
		ItensPerPage int
	}
	tests := []struct {
		name   string
		args   args
		length int
		pages  int
	}{
		{
			name: "find-by-name",
			args: args{
				search:       map[string]interface{}{"name": uuid.NewString()},
				ItensPerPage: 10,
			},
			length: 10,
			pages:  1,
		},
		{
			name: "find-by-surname",
			args: args{
				search:       map[string]interface{}{"surname": uuid.NewString()},
				ItensPerPage: 20,
			},
			length: 20,
			pages:  2,
		},
		// {
		// 	name: "find-by-email",
		// 	args: args{
		// 		search:       map[string]interface{}{"email": uuid.NewString()},
		// 		ItensPerPage: 10,
		// 	},
		// 	length: 1,
		// 	pages:  1,
		// },
	}

	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			db.InitDatabase()

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					for i := 0; i < tt.length; i++ {
						var name, surname, email string
						var active bool

						if tt.args.search["name"] != nil {
							name = tt.args.search["name"].(string)
						} else {
							name = uuid.NewString()
						}

						if tt.args.search["surname"] != nil {
							surname = tt.args.search["surname"].(string)
						} else {
							surname = uuid.NewString()
						}

						if tt.args.search["email"] != nil {
							email = tt.args.search["email"].(string)
						} else {
							email = fmt.Sprintf("%s@test.go", uuid.NewString())
						}

						if tt.args.search["active"] != nil {
							active = tt.args.search["active"].(bool)
						} else {
							active = false
						}

						client := models.Client{
							Name:    &name,
							Surname: &surname,
							Email:   &email,
							Active:  active,
						}

						db.GetInstance().Create(&client)
					}

					got := List(tt.args.search)

					if len(got) != tt.length {
						t.Errorf("List() = %v, want %v", len(got), tt.length)
					}
				})
			}
		})
	}
}
