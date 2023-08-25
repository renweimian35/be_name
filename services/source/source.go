package source

import (
	"be-name/services/common/book"
	"be-name/services/common/result"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const sourcePath = "config/source.json"
const strokePath = "config/utf8_Stroke.txt"
const wordPath = "config/word.json"
const wordGroupPath = "config/ci.json"
const readWordPath = "config/read.txt"

var sourceMap map[int][]book.Content
var strokeMap = make(map[string]int) // 汉字笔画
var wordMap = map[string]result.WordExplain{}
var wordGroupMap = map[string]string{}
var readWord = ""

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

func WordStrokeCount(s string) int {
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

func LoadSingleWordForPath(path string) error {
	w, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return loadSingleWord(w)
}
func LoadSingleWord() error {
	w, err := os.ReadFile(wordPath)
	if err != nil {
		return err
	}
	return loadSingleWord(w)
}

func loadSingleWord(w []byte) error {
	var words []result.WordExplain
	if err := json.Unmarshal(w, &words); err != nil {
		return err
	}
	for _, item := range words {
		wordMap[item.Word] = item
	}
	return nil
}

func WordExplain(word string) result.WordExplain {
	return wordMap[word]
}

type wordGroup struct {
	Ci          string `json:"ci"`
	Explanation string `json:"explanation"`
}

// LoadWordGroup 整个的词组
func LoadWordGroup() error {
	w, err := os.ReadFile(wordGroupPath)
	if err != nil {
		return err
	}
	var wg []wordGroup
	if err = json.Unmarshal(w, &wg); err != nil {
		return err
	}
	for _, item := range wg {
		wordGroupMap[item.Ci] = item.Explanation
	}
	return nil
}

func WordGroupExplain(wg string) (string, bool) {
	explain, ok := wordGroupMap[wg]
	return explain, ok
}

// LoadReadWord 看过的名字就不要随机到了
func LoadReadWord() error {
	p, err := os.ReadFile(readWordPath)
	if err != nil {
		return err
	}
	readWord = string(p)
	return nil
}

func IsReadWord(word string) bool {
	readArr := strings.Split(readWord, ",")
	for _, r := range readArr {
		if r == word {
			return true
		}
	}
	return false
}

// ResetReadWord 把看过的名字，写入文件末尾
func ResetReadWord(readWord string) {
	file, err := os.OpenFile(readWordPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(readWord + "\n")
	if err != nil {
		return
	}
}
