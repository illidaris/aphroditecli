package exporter

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/illidaris/aphrodite/pkg/convert"
)

func FmtCsv(data any, pretty bool) {
	rows := ConvertToRows(data)
	if len(rows) == 0 {
		return
	}
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for _, row := range rows {
		_ = wr.Write(row)
	}
	wr.Flush()
	_ = os.WriteFile(fmt.Sprintf("%v.csv", convert.TimeNumber(time.Now())), b.Bytes(), os.ModePerm)
}
