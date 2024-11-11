package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	t.Run("database connection is fine", func(t *testing.T) {

		mock, db := Mock()

		mock.ExpectQuery("SELECT count\\(\\*\\) FROM information_schema\\.tables .*").
			WithArgs("tasks", "BASE TABLE").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectExec(`CREATE TABLE "tasks" .*`).
			WillReturnResult(sqlmock.NewResult(1, 1))

		ConnectDB(db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("query expectation never met, %q", err.Error())
		}

	})
}
