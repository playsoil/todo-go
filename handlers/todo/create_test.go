package todo

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database"
	"github.com/playsoil/todo-go/models"
)

func TestCreateHandler(t *testing.T) {
	app := fiber.New()
	app.Post("/todo", CreateHandler)

	t.Run("body must be valid json", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/todo", nil)
		resp, _ := app.Test(req, 1)
		if resp.StatusCode != 400 {
			t.Errorf("got %d status wanted %d", resp.StatusCode, 400)
		}
	})

	t.Run("title is required", func(t *testing.T) {
		dummy := struct{ DummyField string }{
			DummyField: "dummy data is here",
		}
		marshaled, err := json.Marshal(dummy)
		if err != nil {
			t.Fatalf("impossible to marshall task: %s", err)
		}
		req := httptest.NewRequest("POST", "/todo", bytes.NewReader(marshaled))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 1)
		if resp.StatusCode != 422 {
			t.Errorf("got %d status wanted %d", resp.StatusCode, 422)
		}
	})

	t.Run("can create title", func(t *testing.T) {
		mock, db := database.Mock()

		database.ConnectDB(db)

		task := models.Task{
			Title: "new task",
		}
		marshaled, err := json.Marshal(task)
		if err != nil {
			t.Fatalf("impossible to marshall task: %s", err)
		}
		req := httptest.NewRequest("POST", "/todo", bytes.NewReader(marshaled))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 1)
		if resp.StatusCode != 200 {
			t.Errorf("got %d status wanted %d", resp.StatusCode, 200)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("query expectation never met, %q", err.Error())
		}
	})
}