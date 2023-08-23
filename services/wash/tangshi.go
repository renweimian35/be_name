package wash

import (
	"be-name/services/common/book"
	"fmt"
)

type tangShi struct {
	Title        string         `json:"title"`
	FirstContent []firstContent `json:"content"`
}

type firstContent struct {
	T        string          `json:"type"`
	Contents []secondContent `json:"content"`
}

type secondContent struct {
	Chapter    string   `json:"chapter"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
}

func (t *tangShi) wash(path, name string) ([]book.Content, error) {
	var ts tangShi
	if err := resolve(path, &ts); err != nil {
		return nil, err
	}
	var bookContent []book.Content
	for _, firstItem := range ts.FirstContent {
		for _, item := range firstItem.Contents {
			bookContent = append(bookContent, book.Content{
				Title:      fmt.Sprintf("%s-%s", name, item.Chapter),
				Author:     item.Author,
				Paragraphs: item.Paragraphs,
			})
		}
	}
	return bookContent, nil
}
