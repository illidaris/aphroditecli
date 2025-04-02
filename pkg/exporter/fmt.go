package exporter

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
)

func FmtTable(data any, pretty bool) {
	rows := ConvertToRows(data)
	if len(rows) == 0 {
		return
	}
	table := tablewriter.NewWriter(os.Stdout) // 输出到标准输出（控制台）
	table.SetHeader(rows[0])
	for _, v := range rows[1:] { // 跳过表头行，只添加数据行
		table.Append(v)
	}
	table.SetBorder(false)                     // 设置不显示边框（如果你想要的话）
	table.SetAlignment(tablewriter.ALIGN_LEFT) // 设置对齐方式
	table.Render()                             // 渲染表格到标准输
}

func FmtJson(data any, pretty bool) {
	bs, err := json.Marshal(data)
	if err != nil || len(bs) == 0 {
		return
	}
	jsonStr := string(bs)
	println(PrettyString(jsonStr, pretty))
}

func ConvertToRows(data any) [][]string {
	switch data := data.(type) {
	case [][]string:
		return data
	case []map[string]any:
		heads := []string{}
		rows := [][]string{}
		for index, kv := range data {
			// 构造列头
			if index == 0 {
				for k := range kv {
					heads = append(heads, k)
				}
				rows = append(rows, heads)
			}
			row := []string{}
			for _, head := range heads {
				row = append(row, cast.ToString(kv[head]))
			}
			rows = append(rows, row)
		}
		return rows
	default:
		return [][]string{}
	}
}

func PrettyString(str string, pretty bool) string {
	if !pretty {
		return str
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str
	}
	return prettyJSON.String()
}
