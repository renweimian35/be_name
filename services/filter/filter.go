package filter

import (
	"be-name/services/common/book"
	"be-name/services/common/result"
	"be-name/services/explain"
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

const beNameCount = 10 //随机十个名字

func BeName(books []book.Content, surname string, count int) {
	rand.NewSource(time.Now().UnixNano())
	var results []result.Info
	for i := 0; i < beNameCount; i++ {
		randomIndex := rand.Intn(len(books))
		re := filterNameByBook(books[randomIndex], surname, count)
		if len(re.Name) != 0 {
			results = append(results, re)
		}
	}

	for _, item := range results {
		if len(item.Name) != 0 {
			fmt.Printf(" %s(%s),出自 %s 的 %s ：\n", item.Name, item.PinYin, item.Author, item.Title)
			fmt.Println(item.Source)
			fmt.Println("---------------------------------------------------------------")
		}
	}
}

// book:传入参数为整本书的内容,包含名字、作者、内容(内容实现处理好，为数组，分段落章节)
// 处理方式：1：随机到某个段落，在取出的章节中再随机出一个汉字
// 2：判断该汉字是否符合要求，如果符合则随机第二个字
// 循序次数，还是找不到则退出
func filterNameByBook(book book.Content, surname string, count int) result.Info {
	info := result.Info{}
	firstWordRule := nextWordRule(surname)
	surnameWordInfo := explain.NewExplain(surname)
	tryCount := 0
	for {
		if tryCount == 20 {
			break
		}
		rand.NewSource(time.Now().UnixNano())
		if len(book.Paragraphs) < 3 {
			break
		}
		paraIndex := rand.Intn(len(book.Paragraphs) - 1)
		firstPara := book.Paragraphs[paraIndex]
		firstWord, err := randString(firstPara)
		if err != nil {
			break
		}
		firstWordInfo, isMatch := wordIsMatch(firstWord, firstWordRule)
		if !isMatch {
			tryCount++
			continue
		}

		secondPara := book.Paragraphs[paraIndex+1]
		secondWord, err := randString(secondPara)
		if err != nil {
			break
		}
		secondWordRule := nextWordRule(firstWord)
		secondWordInfo, isMatch := wordIsMatch(secondWord, secondWordRule)
		if !isMatch {
			tryCount++
			continue
		}
		info.Name = fmt.Sprintf("%s%s%s", surname, firstWord, secondWord)
		info.PinYin = fmt.Sprintf("%s%s%s", surnameWordInfo.Pinyin, firstWordInfo.Pinyin, secondWordInfo.Pinyin)
		info.Stroke = surnameWordInfo.Stroke + firstWordInfo.Stroke + secondWordInfo.Stroke
		info.Author = book.Author
		info.Title = book.Title
		info.Source = fmt.Sprintf("%s\n%s", firstPara, secondPara)
		break
	}
	return info
}

type rule struct {
	passYunMu string //不能为的韵母
	pingZe    int    //平仄:音调1、2声为平，3、4声为仄 ，当前定义：1为平，2为仄
}

// 当前与规则是否匹配
func wordIsMatch(word string, r rule) (explain.Chinese, bool) {
	wordInfo := explain.NewExplain(word)
	if word == "兮" ||
		word == "曰" ||
		wordInfo.Stroke >= 10 ||
		wordInfo.IsManyRead == true ||
		wordInfo.PingZe != r.pingZe ||
		wordInfo.YunMu == r.passYunMu {
		return wordInfo, false
	}
	return wordInfo, true
}

// 寻找下一个字的规则
func nextWordRule(word string) rule {
	info := explain.NewExplain(word)
	pinZe := 1
	if info.PingZe == 1 {
		pinZe = 2
	}
	return rule{
		passYunMu: info.YunMu,
		pingZe:    pinZe,
	}
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
