package encrypts

import (
	"net/url"

	"github.com/illidaris/aphroditecli/pkg/exporter"
)

func UrlEncode(args ...string) [][]string {
	rows := [][]string{
		{"Raw", "Encode"},
	}
	for _, v := range args {
		row := []string{
			v,
			url.QueryEscape(v),
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}

func UrlDecode(args ...string) [][]string {
	rows := [][]string{
		{"Raw", "Decode"},
	}
	for _, v := range args {
		decodeStr, _ := url.QueryUnescape(v)
		row := []string{
			v,
			decodeStr,
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}
