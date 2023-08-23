package main

import (
	"be-name/services/filter"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	serNum    string
	nameCount int
)

func main() {
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

	source, err := filter.SelectSource(serNum)
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}
	filter.BeName(source, nameCount)

}

func init() {
	err := filter.LoadSource()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&serNum, "num", "", "清选择名字来源读物: 1:楚辞，2：诗经，3：唐诗，4：宋词，5：论语，6：曹操诗集。多个以;分开，如：1;2;3")
	flag.IntVar(&nameCount, "count", 0, "请输入几字姓名，如张三则是2，独孤求败则是3")
	flag.Parse()

}
