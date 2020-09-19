package model

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestUser_IsConfirmed(t *testing.T) {

	unconfirmedUser, _ := NewUser("email", "password")
	confirmedUser, _ := NewUser("email", "password")
	confirmedUser.Confirm()

	tests := []struct {
		name string
		u    *User
		want bool
	}{
		{
			name: "User is confirmed",
			u:    confirmedUser,
			want: true,
		},
		{
			name: "User is not confirmed",
			u:    unconfirmedUser,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.IsConfirmed(); got != tt.want {
				t.Errorf("User.IsConfirmed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Authenticate(t *testing.T) {
	password1, err := hashPassword("password1")
	if err != nil {
		panic("Could not hash password")
	}

	type args struct {
		password string
	}
	tests := []struct {
		name string
		u    *User
		args args
		want bool
	}{
		{
			name: "password is valid",
			u: &User{
				EncryptedPassword: password1,
			},
			args: args{password: "password1"},
			want: true,
		},
		{
			name: "password is valid",
			u: &User{
				EncryptedPassword: password1,
			},
			args: args{password: "wrongPassword"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Authenticate(tt.args.password); got != tt.want {
				t.Errorf("User.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty string as password",
			args:    args{password: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Confirm(t *testing.T) {
	tests := []struct {
		name string
		u    *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.u.Confirm()
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "returns a new user",
			args:    args{email: "test@mail.com", password: "password"},
			wantErr: false,
		},
		{
			name:    "fails due to empty email",
			args:    args{email: "", password: "password"},
			wantErr: true,
		},
		{
			name:    "fails due to empty password",
			args:    args{email: "test@mail.com", password: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {

				if got.ID == uuid.Nil {
					t.Errorf("NewUser().ID = %v, want non nil id", got)
				}
				if !(time.Since(got.CreatedAt) > 0) {
					t.Errorf("NewUser().CreatedAt = %v, want timestamp in the past", got)
				}
				if !(time.Since(got.UpdatedAt) > 0) {
					t.Errorf("NewUser().UpdatedAt = %v, want timestamp in the past", got)
				}
				if got.Email != tt.args.email {
					t.Errorf("User.Email = %v, want %v", got, tt.args.email)
				}
				if !got.Authenticate(tt.args.password) {
					t.Errorf("User.EncryptedPassword = %v, did not match hash", got)
				}

				if got.InvitedAt != nil {
					t.Errorf("NewUser().InvitedAt = %v, want nil", got)
				}
				if got.ConfirmedAt != nil {
					t.Errorf("NewUser().ConfirmedAt = %v, want nil", got)
				}
				if got.LastSignInAt != nil {
					t.Errorf("NewUser().LastSignInAt = %v, want nil", got)
				}
				if len(got.Groups) != 0 {
					t.Errorf("NewUser().Groups = %v, want []Group{}", got)
				}
			}
		})
	}
}
