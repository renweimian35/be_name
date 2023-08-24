package source

import (
	"be-name/services/common/book"
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

const sourcePath = "config/source.json"
const strokePath = "config/utf8_Stroke.txt"

var sourceMap map[int][]book.Content
var strokeMap = make(map[string]int) // 汉字笔画

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
	serNumList := strings.Split(serNum, ",")
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

func GetStrokeCount(s string) int {
	for _, r := range s {
		c, ok := strokeMap[string(r)]
		if !ok {
			return 0
		} else {
			return c
		}
		break
	}
	return 0
}

func LoadStrokeForPath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return loadStroke(f)
}

func LoadStroke() error {
	f, err := os.Open(strokePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return loadStroke(f)
}
func loadStroke(f io.Reader) error {
	var line int = 1
	input := bufio.NewScanner(f)
	for input.Scan() {
		ss := strings.Split(input.Text(), " ")
		c, err := strconv.Atoi(ss[2])
		if err != nil {
			return err
		}
		strokeMap[ss[1]] = c
		if line == 20964 {
			break
		} else {
			line++
		}
	}
	return nil
}
