package merge

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	dir := "D:\\work\\data" // 指定要遍历的目录路径
	f, err := os.Create("D:\\work\\result.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	err = fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // 处理错误情况，例如权限问题等。返回错误将停止遍历。
		}
		if !strings.Contains(path, "csv") {
			return nil
		}
		fullPath := filepath.Join(dir, path) // 构造完整的路径名（如果需要的话）
		bs, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		f.Write(bs)
		return nil // 继续遍历。返回nil表示无错误。
	})
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dir, err) // 处理遍历过程中遇到的任何错误。
	}

}
