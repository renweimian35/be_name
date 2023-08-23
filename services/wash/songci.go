package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type songCi struct {
	Rhythmic   string   `json:"rhythmic"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
}

func (c *songCi) wash(path, name string) ([]book.Content, error) {
	var sc []songCi
	if err := resolve(path, &sc); err != nil {
		return nil, err
	}
	var contents []book.Content
	for _, item := range sc {
		contents = append(contents, book.Content{
			Title:      fmt.Sprintf("%s-%s", name, item.Rhythmic),
			Author:     item.Author,
			Paragraphs: item.Paragraphs,
		})
	}
	return contents, nil
}
