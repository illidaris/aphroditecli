package importer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/illidaris/aphrodite/pkg/imex"
	acFile "github.com/illidaris/aphrodite/pkg/io/file"
	"github.com/xuri/excelize/v2"

	fileex "github.com/illidaris/file/path"
)

func ParseFile(file string, opts ...imex.ImExOptionFunc[string]) ([][]string, error) {
	raws := [][]string{}
	b, err := fileex.ExistOrNot(file)
	if err != nil {
		return raws, err
	}
	if !b {
		return raws, fmt.Errorf("%v no exist", file)
	}
	keys := strings.Split(file, ".")
	l := len(keys)
	if l < 1 {
		return raws, fmt.Errorf("%v is error", file)
	}
	f, err := os.Open(file)
	if err != nil {
		return raws, err
	}
	defer f.Close()
	suffix := ""
	if l > 1 {
		suffix = keys[l-1]
	}
	switch suffix {
	case "", "txt":
		bs, err := io.ReadAll(f)
		if err != nil {
			return raws, err
		}
		rawStr := strings.ReplaceAll(string(bs), "\r\n", "\n")
		for _, v := range strings.Split(rawStr, "\n") {
			raws = append(raws, []string{v})
		}
		return raws, err
	case "csv":
		return acFile.CsvImporter(nil)(f)
	case "xlsx", "xls":
		excelReader, err := excelize.OpenReader(f)
		if err != nil {
			return raws, err
		}
		return excelReader.GetRows("Sheet1")
	default:
		return raws, fmt.Errorf("%v is error", file)
	}
}
