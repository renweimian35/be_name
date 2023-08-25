package filter

import (
	"be-name/services/common/book"
	"be-name/services/common/result"
	"be-name/services/explain"
	"be-name/services/source"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const beNameCount = 5 //随机十个名字

func BeName(books []book.Content, surname string, count int) {
	rand.NewSource(time.Now().UnixNano())
	var results []result.Info
	maxNameCount := 10
	nameCount := 0
	for {
		if nameCount == maxNameCount {
			break
		}
		randomIndex := rand.Intn(len(books))
		re := filterNameByBook(books[randomIndex], surname, count)
		if len(re.Name) != 0 {
			nameCount++
			results = append(results, re)
		}
	}

	names := ""
	for _, item := range results {
		if len(item.Name) != 0 {
			names += item.GroupWord + ","
			fmt.Printf(" %s(%s),出自 %s 的 %s ：\n", item.Name, item.PinYin, item.Author, item.Title)
			if len(item.Source) != 0 {
				fmt.Println(item.Source)
			}
			fmt.Printf("其中 %s 的含义为：%s \n", item.GroupWord, item.GroupExplain)
			/*for _, explain := range item.SingleExplain {
				fmt.Printf("其中 %s 的含义为 %s\n", explain.Word, explain.Explanation)
			}*/
			fmt.Println("---------------------------------------------------------------")
		}
	}
	source.ResetReadWord(names)
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
		firstPara = strings.Trim(firstPara, " ")
		firstWord, err := randString(firstPara)
		if err != nil {
			tryCount++
			continue
		}
		firstWordInfo, isMatch := wordIsMatch(firstWord, firstWordRule)
		if !isMatch {
			tryCount++
			continue
		}
		// 如果当前段落大于20个字，一个汉字算3个字符。，则从当前段落取第二个字，否则去下一个段落
		secondPara := ""
		if len(firstPara) > 60 || tryCount < 5 {
			secondPara = firstPara
		} else {
			secondPara = book.Paragraphs[paraIndex+1]
		}
		secondWord, err := randString(secondPara)
		if err != nil {
			tryCount++
			continue
		}
		secondWordRule := nextWordRule(firstWord)
		secondWordInfo, isMatch := wordIsMatch(secondWord, secondWordRule)
		if !isMatch {
			tryCount++
			continue
		}
		nameGroup := fmt.Sprintf("%s%s", firstWord, secondWord)
		if isRead := source.IsReadWord(nameGroup); isRead {
			tryCount++
			continue
		}
		groupExplain, ok := source.WordGroupExplain(nameGroup)
		if !ok {
			tryCount++
			continue
		}
		info.Name = fmt.Sprintf("%s%s%s", surname, firstWord, secondWord)
		info.PinYin = fmt.Sprintf("%s%s%s", surnameWordInfo.Pinyin, firstWordInfo.Pinyin, secondWordInfo.Pinyin)
		info.Stroke = surnameWordInfo.Stroke + firstWordInfo.Stroke + secondWordInfo.Stroke
		info.Author = book.Author
		info.Title = book.Title
		info.GroupExplain = groupExplain
		info.GroupWord = nameGroup
		if firstPara == secondPara {
			info.Source = fmt.Sprintf("%s\n", firstPara)
		} else {
			info.Source = fmt.Sprintf("%s\n%s", firstPara, secondPara)
		}
		info.SingleExplain = append(info.SingleExplain, firstWordInfo.Explain, secondWordInfo.Explain)
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
	if isWordBlack(word) ||
		wordInfo.Stroke >= 10 ||
		wordInfo.IsManyRead == true ||
		wordInfo.PingZe != r.pingZe ||
		wordInfo.YunMu == r.passYunMu {
		return wordInfo, false
	}
	return wordInfo, true
}

func isWordBlack(word string) bool {
	black := []string{"兮", "曰", "日", "死", "我", "玉", "人", "帝", "天", "小", "四", "星", "水", "成", "下", "门", "云", "城", "残", "二", "审", "夜", "乱",
		"老", "旧", "国", "生", "龙", "未", "命", "官", "在", "五", "文", "凤", "杀", "坟", "克", "吊", "殄", "", "", "", "", "", "", "", "", "", ""}
	for _, w := range black {
		if w == word {
			return true
		}
	}
	return false
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
	hanZiRunes := []rune(para)
	if len(hanZiRunes) == 0 {
		return "", errors.New("error")
	}
	// 生成随机索引
	randomIndex := rand.Intn(len(hanZiRunes))
	// 根据随机索引获取对应的汉字
	randomHanZi := string(hanZiRunes[randomIndex])
	return randomHanZi, nil
}

// 定义正则表达式，匹配除汉字之外的字符
// 使用正则表达式替换非汉字字符为空字符串
func removeNonChineseCharacters(input string) string {
	nonChinesePattern := "[^\\p{Han}]"
	re := regexp.MustCompile(nonChinesePattern)
	return re.ReplaceAllString(input, "")
}
