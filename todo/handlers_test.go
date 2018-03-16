package todo

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
	"github.com/mcaparna/GoCodingChallenge/todo"
)

var (
    server   *httptest.Server
    usersUrl string
)

func init() {
    server = httptest.NewServer(todo.Create()) //Creating new server with the user handlers

    usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint
}

func TestCreateTodo(t *testing.T) {
    userJson := `{"status": "Closed", "title": "Clean Home"}`

    reader := strings.NewReader(userJson) //Convert string to reader

    request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
    }
}