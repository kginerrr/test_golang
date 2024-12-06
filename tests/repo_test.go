package tests

import (
	"fibertesttask/internal/model"
	"fibertesttask/internal/repository"
	"testing"

	"database/sql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("error db create %v", err)
	}

	createTableQuery := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	if _, err := db.Exec(createTableQuery); err != nil {
		t.Fatalf("error %v", err)
	}

	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := &model.User{Name: "John Doe", Email: "john@example.com"}
	err := repo.Create(user)

	assert.NoError(t, err)
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	user := &model.User{Name: "Jane Doe", Email: "jane@example.com"}
	err := repo.Create(user)
	assert.NoError(t, err)

	result, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
}
