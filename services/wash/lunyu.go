package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type lunYu struct {
	Chapter    string   `json:"chapter"`
	Paragraphs []string `json:"paragraphs"`
}

func (c *lunYu) wash(path, name string) ([]book.Content, error) {
	var ly []lunYu
	if err := resolve(path, &ly); err != nil {
		return nil, err
	}
	var bookContent []book.Content
	for _, item := range ly {
		bookContent = append(bookContent, book.Content{
			Title:      fmt.Sprintf("%s-%s", name, item.Chapter),
			Author:     "孔子",
			Paragraphs: item.Paragraphs,
		})
	}
	return bookContent, nil
}
