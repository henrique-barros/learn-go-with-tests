package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists    = DictionaryErr("the word is already defined")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Update(term, value string) error {
	_, err := d.Search(term)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[term] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(term string) error {
	_, err := d.Search(term)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, term)
		return nil
	default:
		return err
	}
}

func (d Dictionary) Add(term string, value string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = value
		return nil
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
}
