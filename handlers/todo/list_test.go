package todo

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/playsoil/todo-go/database"
)

func TestListHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/todo", ListHandler)

	t.Run("can see the list", func(t *testing.T) {
		mock, db := database.Mock()

		database.ConnectDB(db)

		// Prepare the mock expected rows (3 rows with ID and title columns)
		rows := sqlmock.NewRows([]string{"ID", "title"}).
			AddRow(1, "Task 1").
			AddRow(2, "Task 2").
			AddRow(3, "Task 3")

		// Expect the query to be executed and return the rows
		mock.ExpectQuery("SELECT \\* FROM \"tasks\"").WillReturnRows(rows)

		req := httptest.NewRequest("GET", "/todo", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 1)
		if resp.StatusCode != 200 {
			t.Errorf("got %d status wanted %d", resp.StatusCode, 200)
		}

		// check if all query expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("query expectation never met, %q", err.Error())
		}

		// check if the response is equal to database rows and what we expected
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("there is error %q", err.Error())
		}
		got := string(body)
		want := `{"data":[{"ID":1,"title":"Task 1"},{"ID":2,"title":"Task 2"},{"ID":3,"title":"Task 3"}]}`

		if got != want {
			t.Errorf("list response is not correct, got %q want %q", got, want)
		}
	})
}
