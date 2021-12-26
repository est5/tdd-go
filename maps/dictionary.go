package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it doesnt exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, definition string) error {
	_, exist := d.Search(word)
	switch exist {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return exist
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
