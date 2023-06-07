package admin

import (
	"miniproject2/entities"
	"miniproject2/repositories"
	"reflect"
	"testing"
)

func Test_useCaseAdmin_CreateAdmin(t *testing.T) {
	type fields struct {
		adminRepo repositories.AdminInterfaceRepo
	}
	type args struct {
		admin AdminParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Admin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseAdmin{
				adminRepo: tt.fields.adminRepo,
			}
			got, err := uc.CreateAdmin(tt.args.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
