package http_error

import "testing"

func TestConflictError(t *testing.T) {
	message := "test message"
	conflictError := ConflictError{Message: message}

	if message != conflictError.Error() {
		t.Errorf("got %s posts, wanted %s posts", conflictError.Error(), message)
	}
}
