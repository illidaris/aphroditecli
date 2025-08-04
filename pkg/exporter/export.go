package exporter

import (
	"strings"
	"time"

	"github.com/illidaris/aphrodite/pkg/convert"
)

func ExportX(name string, data any, pretty bool) {
	var (
		fname   string
		fsuffix string
	)
	keys := strings.Split(name, ".")
	if len(keys) < 1 {
		println("未指定导出格式/名称")
		return
	}
	if len(keys) == 1 {
		fname = time.Now().Format(convert.NumberTimeFormat)
		fsuffix = keys[0]
	}
	if l := len(keys); l > 1 {
		fname = strings.Join(keys[:l-1], ".")
		fsuffix = keys[l-1]
	}
	Export(fsuffix, fname, data, pretty)
}

func Export(exptr, name string, data any, pretty bool) {
	switch exptr {
	case "json":
		FmtJson(data, pretty)
	case "csv":
		FmtCsv(name, data, pretty)
	case "excel":
		FmtExcel(name, data, pretty)
	default:
		FmtTable(data, pretty)
	}
}
