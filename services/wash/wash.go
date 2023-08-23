// Package wash ：由于每个源都是不同的格式，需要进行数据清洗达到同一格式，然后进行导出。
// 后续有新的源加入，只需要实现wash接口，然后在pckSource中进行加入即可。
package wash

import (
	"be-name/services/common/book"
	"encoding/json"
	"fmt"
	"os"
)

const sourceName = "source.json"
const (
	basePath                               = "C:\\Users\\44312\\Desktop\\name_source\\"
	chuCiIndex, chuCiPath, chuCiName       = 1, "chuci.json", "楚辞"
	shiJinIndex, shiJinPath, shiJinName    = iota, "shijin.json", "诗经"
	tangShiIndex, tangShiPath, tangShiName = iota, "tangshi.json", "唐诗"
	songCiIndex, songCiPath, songCiName    = iota, "songci.json", "宋词"
	lunYuIndex, lunYuPath, lunYuName       = iota, "lunyu.json", "论语"
	caoCaoIndex, caoCaoPath, caoCaoName    = iota, "caocao.json", "曹操诗集"
)

func AllSource() error {
	var books []book.Info
	allMap := pckSource()
	for serNum, item := range allMap {
		path := fmt.Sprintf("%s%s", basePath, item.path)
		contents, err := item.source.wash(path, item.name)
		if err != nil {
			return nil
		}

		bookInfo := book.Info{
			Book:    item.name,
			SerNum:  serNum,
			Content: contents,
		}
		books = append(books, bookInfo)
	}
	allSource, err := json.Marshal(books)
	if err != nil {
		return err
	}

	if err := os.WriteFile(sourceName, allSource, os.ModeDevice); err != nil {
		return err
	}
	return nil
}

type sourceInfo struct {
	source source
	path   string
	name   string
}

func pckSource() map[int]sourceInfo {
	sourceMap := make(map[int]sourceInfo)

	sourceMap[chuCiIndex] = sourceInfo{
		source: &chuCi{},
		path:   chuCiPath,
		name:   chuCiName,
	}
	sourceMap[shiJinIndex] = sourceInfo{
		source: &shiJin{},
		path:   shiJinPath,
		name:   shiJinName,
	}

	sourceMap[caoCaoIndex] = sourceInfo{
		source: &caoCao{},
		path:   caoCaoPath,
		name:   caoCaoName,
	}
	sourceMap[songCiIndex] = sourceInfo{
		source: &songCi{},
		path:   songCiPath,
		name:   songCiName,
	}
	sourceMap[lunYuIndex] = sourceInfo{
		source: &lunYu{},
		path:   lunYuPath,
		name:   lunYuName,
	}
	sourceMap[tangShiIndex] = sourceInfo{
		source: &tangShi{},
		path:   tangShiPath,
		name:   tangShiName,
	}
	return sourceMap
}
