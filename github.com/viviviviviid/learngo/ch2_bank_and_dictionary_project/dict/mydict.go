package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
// Dictionary["key"] 이런식으로 찾을 수 있겠지만, method 형태로 만들어주려 한다.
func (d Dictionary) Search(word string) (string, error) {
	// go에서는 사전내용을 찾았을때 두가지 값을 반환해준다. [찾은내용, 존재여부]
	value, exists := d[word]
	if exists { // 존재한다면
		return value, nil
	}
	return "", errNotFound
}
