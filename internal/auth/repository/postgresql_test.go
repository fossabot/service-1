package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
"github.com/stretchr/testify/require"
	"github.com/google/uuid"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/perfolio/service/internal/auth/model"
	"go.uber.org/zap"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// TestMain is responsible for setting up a docker container with a postgresql db and connect to it.
func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "latest",
		Env:          []string{"POSTGRES_PASSWORD=mysecretpassword"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5433"},
			},
		},
	}

	// Run the Docker container
	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Exponential retry to connect to database while it is booting
	if err := pool.Retry(func() error {
		databaseConnStr := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
		db, err = gorm.Open(psql.Open(databaseConnStr), &gorm.Config{})
		if err != nil {
			log.Println("Database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}

		// Tests if database is reachable
		_, err = db.DB()
		return err
	}); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	db.AutoMigrate(&model.User{})

	exitCode := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(exitCode)
}

func TestNewPostgres(t *testing.T) {
	type args struct {
		db     *gorm.DB
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		{
			name: "creates a new postgres repository",
			args: args{
				db:     db,
				logger: zap.NewNop(),
			},
			want: &postgres{db: db, logger: zap.NewNop()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostgres(tt.args.db, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgres_CreateUser(t *testing.T) {
	u, _ := model.NewUser("test@email.com", "password")

	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		p       *postgres
		args    args
		wantErr bool
	}{
		{
			name:    "Successfully inserts a new user",
			p:       &postgres{db: db, logger: zap.NewNop()},
			args:    args{context.Background(), *u},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("postgres.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			var foundUser model.User
			err := tt.p.db.First(&foundUser, tt.args.user.ID).Error
			if err != nil {
				t.Errorf("postgres.CreateUser() error = %v", err)
			}
			require.Equal(t, foundUser.ID, tt.args.user.ID, "ID should be equal")
			require.Equal(t, foundUser.Email, tt.args.user.Email, "Email should be equal")
			require.Equal(t, foundUser.EncryptedPassword, tt.args.user.EncryptedPassword, "EncryptedPassword should be equal")
			require.Equal(t, len(foundUser.Groups), 0, "Groups should be empty")

		})
	}
}
func Test_postgres_ChangeEmail(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       uuid.UUID
		newEmail string
	}
	tests := []struct {
		name    string
		p       *postgres
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ChangeEmail(tt.args.ctx, tt.args.id, tt.args.newEmail); (err != nil) != tt.wantErr {
				t.Errorf("postgres.ChangeEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgres_ConfirmEmail(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		p       *postgres
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ConfirmEmail(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("postgres.ConfirmEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postgres_DeleteUser(t *testing.T) {
	user, _ := model.NewUser("test@mail.com", "password")

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		p       *postgres
		args    args
		wantErr bool
	}{
		{
			name:    "User is not longer in the database",
			p:       &postgres{db: db, logger: zap.NewNop()},
			args:    args{ctx: context.TODO(), id: user.ID},
			wantErr: false,
		},{
			name:    "Fails silently if no user was found",
			p:       &postgres{db: db, logger: zap.NewNop()},
			args:    args{ctx: context.TODO(), id: uuid.New()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		_ = db.Create(user)
		var foundUser model.User
		err := tt.p.db.First(&foundUser, user.ID).Error
		require.NoError(t, err)

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("postgres.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			var foundUser model.User
			err := db.First(&foundUser, tt.args.id).Error
			require.Error(t, err, "Should return error")
		})
	}
}
