package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewUserRepository(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "base case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _, err := sqlmock.New()
			if err != nil {
				t.Errorf("unexpected error setting up test db: %s", err)
			}
			r := NewUserRepository(db)
			if r == nil {
				t.Error("expected: non nil result, got: nil")
			}

		})
	}
}

func Test_userRepository_CreateUser(t *testing.T) {
	type args struct {
		username string
		email    string
		password string
	}
	tests := []struct {
		name        string
		args        args
		setupMockDb func(mock sqlmock.Sqlmock)
		wantErr     bool
	}{
		{
			name: "base case",
			args: args{
				username: "CaptainSteve",
				email:    "USARocks@test.com",
				password: "notsecurebecauseImfromthe50s",
			},
			setupMockDb: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Errorf("unexpected error setting up mock db: %s", err)
			}
			defer db.Close()

			tt.setupMockDb(mock)

			u := &userRepository{
				db: db,
			}
			if err := u.CreateUser(tt.args.username, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

//func Test_userRepository_DeleteUser(t *testing.T) {
//	type fields struct {
//		db sql.DB
//	}
//	type args struct {
//		username string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := &userRepository{
//				db: tt.fields.db,
//			}
//			if err := u.DeleteUser(tt.args.username); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_userRepository_GetUser(t *testing.T) {
//	type fields struct {
//		db sql.DB
//	}
//	type args struct {
//		username string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *models.User
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := &userRepository{
//				db: tt.fields.db,
//			}
//			got, err := u.GetUser(tt.args.username)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_userRepository_UpdateUser(t *testing.T) {
//	type fields struct {
//		db sql.DB
//	}
//	type args struct {
//		username string
//		email    string
//		password string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := &userRepository{
//				db: tt.fields.db,
//			}
//			if err := u.UpdateUser(tt.args.username, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
