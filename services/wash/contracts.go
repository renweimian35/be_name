package wash

import "be-name/services/common/book"

type source interface {
	wash(path, name string) ([]book.Content, error)
}
