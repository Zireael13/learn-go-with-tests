package maps

type DictionaryErr string

const (
	NotFoundError = DictionaryErr("Couldn't find that word")
	WordExistsErr = DictionaryErr("Word already exists")
)

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", NotFoundError
	}

	return def, nil
}

func (d Dictionary) Add(hash, val string) error {
	_, err := d.Search(hash)

	switch err {
	case NotFoundError:
		d[hash] = val
	case nil:
		return WordExistsErr
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = def
		return nil
	}

	return err
}
