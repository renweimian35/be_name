package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type chuCi struct {
	Title   string   `json:"title"`
	Section string   `json:"section"`
	Author  string   `json:"author"`
	Content []string `json:"content"`
}

func (c *chuCi) wash(path, name string) ([]book.Content, error) {
	var cc []chuCi
	if err := resolve(path, &cc); err != nil {
		return nil, err
	}
	var contents []book.Content
	for _, item := range cc {
		contents = append(contents, book.Content{
			Title:      fmt.Sprintf("%s-%s-%s", name, item.Section, item.Title),
			Author:     item.Author,
			Paragraphs: item.Content,
		})
	}
	return contents, nil
}
