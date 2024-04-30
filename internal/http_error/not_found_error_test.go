package http_error

import "testing"

func TestNotFoundError(t *testing.T) {
	modelName := "Test"
	notFoundError := NotFoundError{ModelName: modelName}

	want := "Test not found."
	got := notFoundError.Error()

	if want != got {
		t.Errorf("got %s posts, wanted %s posts", got, want)
	}
}
