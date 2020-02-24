package model

import "testing"

func TestNewPetInstance(t *testing.T) {
	_, err := NewPetInstance("sushi", "This is a description.")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewPetInstance("h", "")
	if err != errorShortName {
		t.Fatalf("expected [%s], got = [%s]", errorShortName, err)
	}

	_, err = NewPetInstance("123456789012345678901234567890123", "")
	if err != errorLongName {
		t.Fatalf("expected [%s], got = [%s]", errorLongName, err)
	}
}

func TestNewPetInstanceWithId(t *testing.T) {
	_, err := NewPetInstanceWithId("1", "sushi", "This is a description.")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewPetInstanceWithId("", "", "")
	if err != errorShortId {
		t.Fatalf("expected [%s], got = [%s]", errorShortId, err)
	}

	tooLong := "0"
	for i := 1; i < 129; i++ {
		tooLong += string(i)
	}

	_, err = NewPetInstanceWithId(tooLong, "", "")
	if err != errorLongId {
		t.Fatalf("expected [%s], got = [%s]", errorLongId, err)
	}
}
