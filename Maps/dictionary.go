package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist to update")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	defination, ok := dictionary[key]
	if !ok {
		return "", ErrWordNotFound
	}
	return defination, nil
}

func (dictionary Dictionary) Add(key string, val string) error {

	_, err := dictionary.Search(key)

	switch err {
	case ErrWordNotFound:
		dictionary[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(key string, val string) error {

	_, err := dictionary.Search(key)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[key] = val
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}
