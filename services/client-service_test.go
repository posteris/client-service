package services

import (
	"testing"

	"github.com/posteris/client-service/models"
)

func TestCreate(t *testing.T) {
	type args struct {
		client *models.Client
		async  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.client, tt.args.async); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
