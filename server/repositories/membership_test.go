package repositories

import (
	"golang_web_programming/server/model"
	"reflect"
	"testing"
)

func TestMembershipRepository_GetByID(t *testing.T) {
	type fields struct {
		data map[string]model.Membership
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Membership
		wantErr bool
	}{
		{"ID가 1인 멤버십 조회한다",
			fields{
				data: map[string]model.Membership{
					"1": {
						ID:             "1",
						UserName:       "test",
						MembershipType: "naver",
					},
				},
			},
			args{id: "1"},
			&model.Membership{
				ID:             "1",
				UserName:       "test",
				MembershipType: "naver",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MembershipRepository{
				data: tt.fields.data,
			}
			got, err := r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
