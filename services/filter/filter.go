package filter

import (
	"be-name/services/common/book"
	"be-name/services/common/result"
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

const beNameCount = 2 //随机十个名字

func BeName(books []book.Content, count int) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < beNameCount; i++ {
		randomIndex := rand.Intn(len(books))
		filterNameByBook(books[randomIndex], count)
	}
}

func filterNameByBook(book book.Content, count int) result.Info {
	info := result.Info{}
	for {
		rand.NewSource(time.Now().UnixNano())
		if len(book.Paragraphs) < 3 {
			break
		}
		paraIndex := rand.Intn(len(book.Paragraphs) - 1)
		para := book.Paragraphs[paraIndex]
		firstWord, err := randString(para)
		if err != nil {
			break
		}
		nextPara := book.Paragraphs[paraIndex+1]
		secondWord, err := randString(nextPara)
		if err != nil {
			break
		}
		fmt.Printf("阳%s%s   取自:\n", firstWord, secondWord)
		fmt.Println(para)
		fmt.Println(nextPara)
		break
	}
	return info
}

// 随机取出一个汉字
// t:0表示从前半部分，1表示从后半部分，2表示无所谓
func randString(para string) (string, error) {
	rand.NewSource(time.Now().UnixNano())
	para = removeNonChineseCharacters(para)
	paraLen := utf8.RuneCountInString(para)
	// 随机选取一个汉字
	randomIndex := rand.Intn(paraLen)
	var currentIndex int
	for _, char := range para {
		if utf8.RuneLen(char) == 3 {
			if currentIndex == randomIndex {
				return string(char), nil
			}
			currentIndex++
		}
	}
	return "", fmt.Errorf("找不到随机字符")
}

// 定义正则表达式，匹配除汉字之外的字符
// 使用正则表达式替换非汉字字符为空字符串
func removeNonChineseCharacters(input string) string {
	nonChinesePattern := "[^\\p{Han}]"
	re := regexp.MustCompile(nonChinesePattern)
	return re.ReplaceAllString(input, "")
}
