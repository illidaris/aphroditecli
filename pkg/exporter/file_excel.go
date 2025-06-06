package exporter

import (
	"fmt"
	"time"

	"github.com/illidaris/aphrodite/pkg/convert"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

func FmtExcel(name string, data any, pretty bool) {
	rows := ConvertToRows(data)
	if len(rows) == 0 {
		return
	}
	f := excelize.NewFile()
	for rowIndex, row := range rows {
		for colIndex, v := range row {
			cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			_ = f.SetCellValue("Sheet1", cellName, v)
		}
	}
	if name == "" {
		name = cast.ToString(convert.TimeNumber(time.Now()))
	}
	_ = f.SaveAs(fmt.Sprintf("%v.xlsx", name))
}
