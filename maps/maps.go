package maps

type DictionaryErr string

const (
	NotFoundErr     = DictionaryErr("Couldn't find that word")
	WordExistsErr   = DictionaryErr("Word already exists")
	WordNotExistErr = DictionaryErr("Word doesn't exist")
)

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", NotFoundErr
	}

	return def, nil
}

func (d Dictionary) Add(hash, val string) error {
	_, err := d.Search(hash)

	switch err {
	case NotFoundErr:
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

	switch err {
	case NotFoundErr:
		return WordNotExistErr
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
