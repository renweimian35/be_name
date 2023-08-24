package explain

import (
	"be-name/services/source"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExplain(t *testing.T) {
	fmt.Println(os.Getwd())
	// 获取当前绝对路径
	executable, err := os.Getwd()
	if err != nil {
		fmt.Println("获取可执行文件路径时出错：", err)
		return
	}
	// 上级目录（根目录）
	rootDir := filepath.Dir(filepath.Dir(executable))
	// 构建文件路径
	strokePath := filepath.Join(rootDir, "config", "utf8_Stroke.txt")
	source.LoadStrokeForPath(strokePath)
	NewExplain("约")
}
