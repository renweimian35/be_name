package filter

import (
	"be-name/services/common/book"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

const sourcePath = "source.json"

var sourceMap map[int][]book.Content

func LoadSource() error {
	sourceMap = make(map[int][]book.Content)
	s, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	var books []book.Info
	err = json.Unmarshal(s, &books)
	if err != nil {
		return err
	}
	for _, item := range books {
		sourceMap[item.SerNum] = item.Content
	}
	return nil
}
func SelectSource(serNum string) ([]book.Content, error) {
	var bookList []book.Content
	serNumList := strings.Split(serNum, ";")
	for _, ser := range serNumList {
		index, err := strconv.Atoi(ser)
		if err != nil {
			return bookList, err
		}
		content, ok := sourceMap[index]
		if !ok {
			return bookList, errors.New("num input error")
		}
		bookList = append(bookList, content...)
	}
	return bookList, nil
}
