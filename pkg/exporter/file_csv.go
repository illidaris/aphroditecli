package exporter

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/illidaris/aphrodite/pkg/convert"
	"github.com/spf13/cast"
)

func FmtCsv(name string, data any, pretty bool) {
	rows := ConvertToRows(data)
	if len(rows) == 0 {
		return
	}
	b := &bytes.Buffer{}
	b.WriteString("\xEF\xBB\xBF")
	wr := csv.NewWriter(b)
	for _, row := range rows {
		_ = wr.Write(row)
	}
	wr.Flush()
	if name == "" {
		name = cast.ToString(convert.TimeNumber(time.Now()))
	}
	_ = os.WriteFile(fmt.Sprintf("%v.csv", name), b.Bytes(), os.ModePerm)
}
