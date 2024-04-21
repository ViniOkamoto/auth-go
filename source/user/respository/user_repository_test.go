package repository

// import (
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/viniokamoto/go-store/source/user/domain/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func TestFindAll_FailedConnection_ReturnError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnError(gorm.ErrInvalidDB)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.FindAll()
// 	assert.Error(t, err)
// }

// func TestFindAll_Success_ReturnUsers(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password", "role_id"})

// 	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnRows(rows)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.FindAll()
// 	assert.NoError(t, err)
// }

// func TestCreate_FailedConnection_ReturnError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectExec("^INSERT INTO \"users\"").WillReturnError(gorm.ErrInvalidDB)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.Create(models.UserCreateRequest{})
// 	assert.Error(t, err)
// }

// func TestCreate_Success_ReturnUser(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectBegin()
// 	mock.ExpectQuery("^INSERT INTO \"users\" (.+) RETURNING \"id\"").
// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
// 		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
// 	mock.ExpectCommit()

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.Create(models.UserCreateRequest{})
// 	assert.NoError(t, err)
// }

// func TestFindById_FailedConnection_ReturnError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnError(gorm.ErrInvalidDB)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.FindById(1)
// 	assert.Error(t, err)
// }

// func TestFindById_Success_ReturnUser(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "email", "password", "role_id"}).
// 		AddRow(1, "foo", "bar", "foobar", "foo@bar.ca", "foobar", 1)

// 	mock.ExpectQuery("^SELECT (.+) FROM \"users\"").WillReturnRows(rows)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.FindById(1)
// 	assert.NoError(t, err)
// }

// func TestUpdate_FailedConnection_ReturnError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectExec("^UPDATE \"users\"").WillReturnError(gorm.ErrInvalidDB)

// 	repo := UserRepositoryFactory(gdb)

// 	_, err = repo.Update(models.User{}, models.UserUpdateRequest{})
// 	assert.Error(t, err)
// }
// func TestUpdate_Success_ReturnUser(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	// Expectation for the Create method
// 	mock.ExpectBegin()
// 	mock.ExpectQuery("^INSERT INTO \"users\"").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
// 	mock.ExpectCommit()

// 	// Expectation for the Create method
// 	mock.ExpectBegin()
// 	mock.ExpectQuery("^INSERT INTO \"users\"").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "barz", "bar", "foobar", "foobar", "test@test.ca", 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
// 	mock.ExpectCommit()

// 	repo := UserRepositoryFactory(gdb)

// 	user := models.User{
// 		FirstName: "foo",
// 		LastName:  "bar",
// 		Username:  "foobar",
// 		Email:     "foo@bar.ca",
// 		Password:  "foobar",
// 		RoleID:    1,
// 	}
// 	createUserRequest := models.UserCreateRequest{
// 		FirstName:       user.FirstName,
// 		LastName:        user.LastName,
// 		Username:        user.Username,
// 		Email:           user.Email,
// 		Password:        user.Password,
// 		ConfirmPassword: user.Password,
// 		RoleID:          user.RoleID,
// 	}

// 	_, err = repo.Create(createUserRequest)
// 	assert.NoError(t, err)

// 	_, err = repo.Update(
// 		user,
// 		models.UserUpdateRequest{
// 			FirstName: "barz",
// 			LastName:  "bar",
// 			Username:  "foobar",
// 			Email:     "test@test.ca",
// 			RoleID:    1,
// 		},
// 	)

// 	assert.NoError(t, err)
// }

// func TestDelete_FailedConnection_ReturnError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	mock.ExpectExec("^DELETE FROM \"users\"").WillReturnError(gorm.ErrInvalidDB)

// 	repo := UserRepositoryFactory(gdb)

// 	err = repo.Delete(models.User{})
// 	assert.Error(t, err)
// }

// func TestDelete_Success_ReturnNil(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: db})
// 	gdb, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening gorm database", err)
// 	}

// 	// Expectation for the Create method
// 	mock.ExpectBegin()
// 	mock.ExpectQuery("^INSERT INTO \"users\"").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
// 	mock.ExpectCommit()

// 	repo := UserRepositoryFactory(gdb)

// 	user := models.User{
// 		FirstName: "foo",
// 		LastName:  "bar",
// 		Username:  "foobar",
// 		Email:     "foo@bar",
// 		Password:  "foobar",
// 		RoleID:    1,
// 	}
// 	createUserRequest := models.UserCreateRequest{
// 		FirstName:       user.FirstName,
// 		LastName:        user.LastName,
// 		Username:        user.Username,
// 		Email:           user.Email,
// 		Password:        user.Password,
// 		ConfirmPassword: user.Password,
// 		RoleID:          user.RoleID,
// 	}

// 	_, err = repo.Create(createUserRequest)
// 	assert.NoError(t, err)

// 	// Expectation for the Delete method
// 	mock.ExpectBegin()
// 	mock.ExpectExec("^UPDATE \"users\" SET \"deleted_at\"='.*' WHERE \"users\".\"deleted_at\" IS NULL").WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	err = repo.Delete(user)
// 	assert.NoError(t, err)
// }
