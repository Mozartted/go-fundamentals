package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("Get the value successfully", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test", "test2": "section 2"}

		got, _ := dictionary.Search("test")

		want := "this is just a test"
		// assertError(t, error, ErrNotFound)
		assertStrings(t, got, want)
	})

	t.Run("Fail in getting the value", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		if _, error := dictionary.Search("flow"); error != nil {
			assertError(t, error, ErrNotFound)
		}
	})

	t.Run("Add to Dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "condensation flow"}

		dictionary.Push("title", "value")

		got, error := dictionary.Search("title")
		if error != nil {
			t.Fatal(error.Error())
		}

		want := "value"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
