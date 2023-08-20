package use_case

import (
	"ddd_prac/port/out/repository/user"
	"errors"
	"reflect"
	"testing"
)

type mockUserRepositorySuccess struct {
}

type mockUserRepositoryFail struct {
}

func (mur mockUserRepositorySuccess) SearchUser(string, string) (int64, error) {
	return 1, nil
}

func (mur mockUserRepositoryFail) SearchUser(string, string) (int64, error) {
	return 0, errors.New("로그인 실패 ㅠ")
}

func TestUserServiceImpl_Login(t *testing.T) {
	type fields struct {
		repository user.Repository
	}
	type args struct {
		email string
		pw    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantL   *LoginUser
		wantErr bool
	}{
		{
			name:   "success login",
			fields: fields{repository: mockUserRepositorySuccess{}},
			args: args{
				email: "joy@naver.com",
				pw:    "world2",
			},
			wantL:   &LoginUser{1},
			wantErr: false,
		}, {
			name:   "fail login",
			fields: fields{repository: mockUserRepositoryFail{}},
			args: args{
				email: "joy@naver.com",
				pw:    "world2",
			},
			wantL:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServiceImpl{
				repository: tt.fields.repository,
			}
			gotL, err := u.Login(tt.args.email, tt.args.pw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("Login() gotL = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
