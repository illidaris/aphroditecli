package encrypts

import (
	"encoding/base64"

	"github.com/illidaris/aphroditecli/pkg/exporter"
)

func Base64Encode(args ...string) [][]string {
	rows := [][]string{
		{"Raw", "StdEncode", "UrlEncode"},
	}
	for _, v := range args {
		row := []string{
			v,
			base64.StdEncoding.EncodeToString([]byte(v)),
			base64.URLEncoding.EncodeToString([]byte(v)),
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}

func Base64Decode(args ...string) [][]string {
	rows := [][]string{
		{"Raw", "StdDecode", "UrlDecode"},
	}
	f := func(decodefunc func(s string) ([]byte, error), raw string) string {
		v, err := decodefunc(raw)
		if err != nil {
			return ""
		}
		return string(v)
	}
	for _, v := range args {
		row := []string{
			v,
			f(base64.StdEncoding.DecodeString, v),
			f(base64.URLEncoding.DecodeString, v),
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}
