package main

import (
	"be-name/services/filter"
	"be-name/services/source"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	serNum    string
	nameCount int
	surname   string
)

func main() {
	serNum = "5"
	nameCount = 3
	surname = "阳"
	flag.Usage = func() {
		fmt.Println("请输入如下信息：")
		flag.PrintDefaults()
	}
	if nameCount == 0 || nameCount > 4 {
		flag.Usage()
		os.Exit(1)
	}
	if len(serNum) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if len(surname) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	bookSource, err := source.SelectSource(serNum)
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}
	filter.BeName(bookSource, surname, nameCount)

}

func init() {
	err := source.LoadSource()
	if err != nil {
		log.Fatal(err)
	}
	err = source.LoadStroke()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&serNum, "num", "", "清选择名字来源读物: 1:楚辞，2：诗经，3：唐诗，4：宋词，5：论语，6：曹操诗集。多个以,分开，如：1,2,3")
	flag.StringVar(&surname, "surname", "", "请输入姓氏")
	flag.IntVar(&nameCount, "count", 0, "请输入几字姓名，如张三则是2，独孤求败则是3")
	flag.Parse()

}
