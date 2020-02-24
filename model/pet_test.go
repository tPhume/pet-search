package model

import "testing"

func TestNewPetInstance(t *testing.T) {
	_, err := NewPetInstance("sushi", "This is a description.")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewPetInstance("", "")
	if err != errorEmptyName {
		t.Fatalf("expected [%s], got = [%s]", errorEmptyName, err)
	}

	_, err = NewPetInstance("h", "")
	if err != errorShortName {
		t.Fatalf("expected [%s], got = [%s]", errorShortName, err)
	}

	_, err = NewPetInstance("12345678901234567890123456789012", "")
	if err != errorLongName {
		t.Fatalf("expected [%s], got = [%s]", errorLongName, err)
	}
}
