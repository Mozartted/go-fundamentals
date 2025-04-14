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
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")

		want := "this is just a test"
		// assertError(t, error, ErrNotFound)
		assertStrings(t, got, want)
	})

	t.Run("Fail in getting the value", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, error := dictionary.Search("flow")
		assertError(t, error, ErrNotFound)
	})
}
