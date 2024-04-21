package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/viniokamoto/go-store/source/user/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindAll_FailedConnection_ReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnError(gorm.ErrInvalidDB)

	repo := UserRepositoryFactory(gdb)

	_, err = repo.FindAll()
	assert.Error(t, err)
}

func TestFindAll_Success_ReturnUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password", "role_id"})

	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnRows(rows)

	repo := UserRepositoryFactory(gdb)

	_, err = repo.FindAll()
	assert.NoError(t, err)
}

func TestCreate_FailedConnection_ReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	mock.ExpectExec("^INSERT INTO \"users\"").WillReturnError(gorm.ErrInvalidDB)

	repo := UserRepositoryFactory(gdb)

	_, err = repo.Create(models.UserRequest{})
	assert.Error(t, err)
}

func TestCreate_Success_ReturnUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery("^INSERT INTO \"users\" (.+) RETURNING \"id\"").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repo := UserRepositoryFactory(gdb)

	_, err = repo.Create(models.UserRequest{})
	assert.NoError(t, err)
}
