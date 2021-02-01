package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find word")
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("nonexisting word", func(t *testing.T) {
		_, err := dictionary.Search("unk")

		if err == nil {
			t.Fatal("No error returned")
		}

		assertError(t, err, NotFoundError)

	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		hash := "test"
		val := "testertest"

		err := dict.Add(hash, val)

		assertError(t, err, nil)

		assertDefinition(t, dict, hash, val)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new")

		assertError(t, err, WordExistsErr)
		assertDefinition(t, dict, word, def)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"

		firstdef := "hello"

		adderr := dict.Add(word, firstdef)

		seconddef := "world"

		updateerr := dict.Update(word, seconddef)

		assertError(t, adderr, nil)
		assertError(t, updateerr, nil)

		assertDefinition(t, dict, word, seconddef)

	})

}
