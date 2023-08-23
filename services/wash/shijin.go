package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type shiJin struct {
	Title   string   `json:"title"`
	Section string   `json:"section"`
	Content []string `json:"content"`
}

func (c *shiJin) wash(path, name string) ([]book.Content, error) {
	var sj []shiJin
	if err := resolve(path, &sj); err != nil {
		return nil, err
	}

	var contents []book.Content
	for _, item := range sj {
		contents = append(contents, book.Content{
			Title:      fmt.Sprintf("%s-%s-%s", name, item.Title, item.Section),
			Author:     "",
			Paragraphs: item.Content,
		})
	}
	return contents, nil
}
