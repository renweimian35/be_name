package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type caoCao struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
}

func (c *caoCao) wash(path, name string) ([]book.Content, error) {
	var cc []caoCao
	if err := resolve(path, &cc); err != nil {
		return nil, err
	}
	var bookContent []book.Content
	for _, item := range cc {
		bookContent = append(bookContent, book.Content{
			Title:      fmt.Sprintf("%s-%s", name, item.Title),
			Author:     "曹操",
			Paragraphs: item.Paragraphs,
		})
	}
	return bookContent, nil
}
