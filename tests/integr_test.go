package tests

import (
	"bytes"
	"encoding/json"
	"fibertesttask/internal/db"
	"fibertesttask/internal/repository"
	"fibertesttask/internal/service"
	"net/http"
	"net/http/httptest"
	"os/user"
	"testing"

	"fibertesttask/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	db.InitDatabase()

	repo := repository.NewUserRepository(db.DB)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	app := fiber.New()
	app.Post("/users", handler.Create)
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := service.GetAllUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(users)
	})

	return app
}

func TestCreateUserIntegration(t *testing.T) {
	app := setupTestApp()

	payload := map[string]string{
		"name":  "John Doe",
		"email": "john@example.com",
	}
	payloadBytes, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestGetAllUsersIntegration(t *testing.T) {
	app := setupTestApp()

	_, _ = app.Test(httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"name":"Jane Doe","email":"jane@example.com"}`))), -1)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var users []user.User
	json.NewDecoder(resp.Body).Decode(&users)

	assert.NotEmpty(t, users)
	assert.Equal(t, "Jane Doe", users[0].Name)
}
