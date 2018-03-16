package todo

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
	"github.com/mcaparna/GoCodingChallenge/todo"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)


func TestUpdateTodo(t *testing.T) {
    userJson := `{"status": "Closed", "title": "Clean Home"}`

    reader := strings.NewReader(userJson) //Convert string to reader

    request, err := http.NewRequest("PUT", "http://127.0.0.1:8080/todos/1", reader) //Create request with JSON body
	
	res := httptest.NewRecorder()

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM todo WHERE").WithArgs("id","title","status").WillReturnRows(sqlmock.NewRows([]string{"id","title", "status"}).AddRow("1","val1", "val2"))
	
	mock.ExpectPrepare("UPDATE todo SET title = $2, status = $3 WHERE id = $1").ExpectExec().WithArgs("id","title","status").WillReturnResult(sqlmock.NewResult(1, 1))
	
	mock.ExpectClose()
	
	todo.Update(res, request, nil)
	
	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTodo(t *testing.T) {
    userJson := `{"status": "Closed", "title": "Clean Home"}`

    reader := strings.NewReader(userJson) //Convert string to reader

    request, err := http.NewRequest("POST", "http://127.0.0.1:8080/todos", reader) //Create request with JSON body
	
	res := httptest.NewRecorder()

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM todo WHERE").WithArgs("id","title","status").WillReturnRows(sqlmock.NewRows([]string{"id","title", "status"}).AddRow("1","val1", "val2"))
	
	mock.ExpectPrepare("INSERT into todo").ExpectExec().WithArgs("id","title","status").WillReturnResult(sqlmock.NewResult(1, 1))
	
	mock.ExpectClose()
	
	todo.Create(res, request, nil)
	
	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestListTodo(t *testing.T) {
    userJson := `{"status": "Closed", "title": "Clean Home"}`

    reader := strings.NewReader(userJson) //Convert string to reader

    request, err := http.NewRequest("GET", "http://127.0.0.1:8080/todos", reader) //Create request with JSON body
	
	res := httptest.NewRecorder()

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM todo WHERE").WithArgs("id","title","status").WillReturnRows(sqlmock.NewRows([]string{"id","title", "status"}).AddRow("1","val1", "val2"))
	
	mock.ExpectClose()
	
	todo.List(res, request, nil)
	
	if err != nil {
		t.Errorf("Expected no error, but got %s instead", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
