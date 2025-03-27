package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "that's just a test"}
	t.Run("default", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "that's just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("default add", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary = Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing update", func(t *testing.T) {
		word := "test"
		definition := "original value"
		dictionary := Dictionary{word: definition}
		newDefinition := "updated value"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new update", func(t *testing.T) {
		word := "test"
		definition := "original value"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrorNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
