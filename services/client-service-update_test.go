package services

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	conndata "github.com/posteris/ci-database-starter/conn-data"
	"github.com/posteris/client-service/db"
	"github.com/posteris/client-service/models"
	"github.com/posteris/database"
)

func TestUpdate(t *testing.T) {
	//obtains the connection database parameters from the posteris/ci-database-starter
	var databases = conndata.GetTestData()

	//create test cases
	type args struct {
		create models.Client
		update models.Client
		want   models.Client
	}
	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	//execute the all test cases for each database
	for _, dbData := range databases {
		t.Run(dbData.Name, func(t *testing.T) {
			//set environment variables
			os.Setenv(database.DatabaseTypeLabel, dbData.Type)
			os.Setenv(database.DatabaseDsnLabel, dbData.DSN)

			//start database
			db.InitDatabase()

			//create test cases
			var tests []testCase

			//lower-case-test
			lowerCaseOBJ := *createNewClientToUpdate()
			lowerCaseEmail := fmt.Sprintf("%s@test.go", uuid.NewString())
			lowerCase := testCase{
				name: "lower-case",
				args: args{
					create: lowerCaseOBJ,
					update: *createClient(lowerCaseOBJ.ID, "jane", "doo", lowerCaseEmail, false),
					want:   *createClient(lowerCaseOBJ.ID, "jane", "doo", lowerCaseEmail, false),
				},
				wantErr: false,
			}
			tests = append(tests, lowerCase)

			//upper-case-test
			upperCaseOBJ := *createNewClientToUpdate()
			upperCaseEmail := fmt.Sprintf("%s@TEST.GO", uuid.NewString())
			upperCase := testCase{
				name: "upper-case",
				args: args{
					create: upperCaseOBJ,
					update: *createClient(upperCaseOBJ.ID, "JANE", "DOO", upperCaseEmail, false),
					want:   *createClient(upperCaseOBJ.ID, "jane", "doo", strings.ToLower(upperCaseEmail), false),
				},
				wantErr: false,
			}
			tests = append(tests, upperCase)

			//no-name-test
			noNameOBJ := *createNewClientToUpdate()
			noNameEmail := fmt.Sprintf("%s@test.go", uuid.NewString())
			noName := testCase{
				name: "no-name",
				args: args{
					create: noNameOBJ,
					update: *createClient(noNameOBJ.ID, "", "doo", noNameEmail, false),
					want:   *createClient(noNameOBJ.ID, *noNameOBJ.Name, "doo", noNameEmail, false),
				},
				wantErr: false,
			}
			tests = append(tests, noName)

			//no-surname-test
			noSurnameOBJ := *createNewClientToUpdate()
			noSurnameEmail := fmt.Sprintf("%s@test.go", uuid.NewString())
			noSurname := testCase{
				name: "no-surname",
				args: args{
					create: noSurnameOBJ,
					update: *createClient(noSurnameOBJ.ID, "jane", "", noSurnameEmail, false),
					want:   *createClient(noSurnameOBJ.ID, "jane", *noSurnameOBJ.Surname, noSurnameEmail, false),
				},
				wantErr: false,
			}
			tests = append(tests, noSurname)

			//no-email-test
			noEmailOBJ := *createNewClientToUpdate()
			noEmail := testCase{
				name: "no-surname",
				args: args{
					create: noEmailOBJ,
					update: *createClient(noEmailOBJ.ID, "jane", "doo", "", false),
					want:   *createClient(noEmailOBJ.ID, "jane", "doo", *noEmailOBJ.Email, false),
				},
				wantErr: false,
			}
			tests = append(tests, noEmail)

			//activated
			activatedOBJ := *createNewClientToUpdate()
			activated := testCase{
				name: "activated",
				args: args{
					create: activatedOBJ,
					update: *createClient(activatedOBJ.ID, "jane", "doo", *activatedOBJ.Email, true),
					want:   *createClient(activatedOBJ.ID, "jane", "doo", *activatedOBJ.Email, false),
				},
				wantErr: false,
			}
			tests = append(tests, activated)

			//email-duplicated
			dupOBJ := *createNewClientToUpdate()
			dupOBJ2 := *createNewClientToUpdate()
			dup := testCase{
				name: "duplicated-email",
				args: args{
					create: dupOBJ,
					update: *createClient(dupOBJ.ID, "", "", *dupOBJ2.Email, false),
					want:   *createClient(dupOBJ.ID, *dupOBJ.Name, *dupOBJ.Surname, *dupOBJ.Email, false),
				},
				wantErr: true,
			}
			tests = append(tests, dup)

			//execute all test cases
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					//update registry
					if err := Update(&tt.args.update); (err != nil) != tt.wantErr {
						t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					}

					//find registry in the database
					got := GetById(tt.args.create.ID)

					//check if update was donne
					if !got.Equals(&tt.args.want) {
						t.Errorf("Update not work! got %v, but want %v", got.ToString(), tt.args.want.ToString())
					}
				})
			}
		})
	}
}
