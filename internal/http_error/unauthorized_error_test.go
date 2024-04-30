package http_error

import "testing"

func TestUnauthorizedErrorError(t *testing.T) {
	unauthorizedError := UnauthorizedError{}

	want := "Unauthorized."
	got := unauthorizedError.Error()

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
