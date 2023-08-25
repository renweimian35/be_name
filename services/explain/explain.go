// Package explain 汉字的拆解，如拼音，笔画，平仄，声律等
package explain

import (
	"be-name/services/common/result"
	"be-name/services/source"
	"github.com/mozillazg/go-pinyin"
)

type Chinese struct {
	Word       string //单汉字
	Pinyin     string
	Stroke     int                //笔画
	YunMu      string             //韵母
	PingZe     int                //平仄:音调1、2声为平，3、4声为仄 ，当前定义：1为平，2为仄
	IsManyRead bool               //是否为多音字
	Explain    result.WordExplain //具体的解释
}

func NewExplain(word string) Chinese {
	ch := Chinese{Word: word}
	ch.pkg()
	return ch
}

func (c *Chinese) pkg() {
	py := pinyin.Pinyin(c.Word, pinyin.Args{
		Style:     pinyin.Tone,
		Heteronym: true,
	})
	if len(py) != 0 {
		pyArray := py[0] // 支持多个汉字，取第一个
		c.Pinyin = pyArray[0]
		if len(pyArray) != 1 {
			c.IsManyRead = true
		}
	}

	c.Stroke = source.WordStrokeCount(c.Word)

	yunMus := pinyin.Pinyin(c.Word, pinyin.Args{
		Style: pinyin.Finals,
	})
	if len(yunMus) != 0 {
		c.YunMu = yunMus[0][0]
	}

	pingZes := pinyin.Pinyin(c.Word, pinyin.Args{
		Style: pinyin.Tone3,
	})
	if len(pingZes) != 0 {
		py := pingZes[0][0]
		pyNum := py[len(py)-1:]
		if pyNum == "3" || pyNum == "4" {
			c.PingZe = 2
		} else {
			c.PingZe = 1
		}
	}
	explain := source.WordExplain(c.Word)
	if explain.Word == c.Word {
		c.Explain = explain
	}
}
