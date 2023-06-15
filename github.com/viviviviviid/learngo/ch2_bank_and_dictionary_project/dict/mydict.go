package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// Dictionary["key"] 이런식으로 찾을 수 있겠지만, method 형태로 만들어주려 한다.
	value, exists := d[word]
	// go에서는 사전내용을 찾았을때 두가지 값을 반환해준다. [찾은내용, 존재여부]
	if exists {
		// 존재한다면
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	// 바로 위 Search 메소드에서는 찾은 값과 에러를 반환한다. // 여기서 우리는 에러가 왔는지만 확인해도 됨
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil: // 검색했을때 에러가 없으므로 아직 존재하지 않음
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Deleten a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
