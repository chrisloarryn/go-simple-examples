package utils

import "testing"

func TestReverse(t *testing.T) {
	var (
		value = "this is a text that will be reversed"
	)

	reversed, err := Reverse(value)
	if err != nil {
		t.Error(err)
	}

	if reversed != "desrever eb lliw taht txet a si siht" {
		t.Errorf("expected %s, got %s", "desrever eb lliw taht txet a si siht", reversed)
	}

	reversed, err = Reverse("")
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != ErrEmpty {
		t.Errorf("expected %s, got %s", ErrEmpty, err.Error())
	}
}
