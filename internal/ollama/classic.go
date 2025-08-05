package ollama

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/illidaris/aphrodite/pkg/convert/table2struct"
	hOllamaBase "github.com/illidaris/aphrodite/pkg/ollama/biz/base"
	hOllamaClassic "github.com/illidaris/aphrodite/pkg/ollama/biz/classic"
	"github.com/illidaris/aphroditecli/pkg/exporter"
	"github.com/illidaris/aphroditecli/pkg/importer"
	fileex "github.com/illidaris/file/path"
	"github.com/spf13/cast"
)

// 输入文件，输出文件标注
func Classic(ctx context.Context, host, model, template, labelFile, categoryFile, out string, args []string) error {
	labels, err := labelsFrmFile(labelFile)
	if err != nil {
		return err
	}
	categories, err := textFrmFile(categoryFile)
	if err != nil {
		return err
	}
	var id int64 = 0
	entries := []*hOllamaBase.Entry{}
	for index, v := range args {
		b, err := fileex.ExistOrNot(v)
		if err != nil || !b {
			entries = append(entries, &hOllamaBase.Entry{
				Id:      atomic.AddInt64(&id, 1),
				Code:    cast.ToString(index),
				Content: v,
			})
			continue
		}
		txts, err := textFrmFile(v)
		if err != nil {
			return err
		}
		for _, txt := range txts {
			entries = append(entries, &hOllamaBase.Entry{
				Id:      atomic.AddInt64(&id, 1),
				Code:    v,
				Content: txt,
			})
		}
	}
	err = hOllamaClassic.ClassicFunc(host, template,
		hOllamaBase.WithModel(model),
		hOllamaBase.WithThink(false),
		hOllamaBase.WithHandle(func(ctx context.Context, e *hOllamaBase.Entry) error {
			fmt.Printf("[%v]结果(%vms )：%v %v \n", e.Id, e.Duration, e.Result, e.Content)
			return nil
		}))(context.Background(), categories, labels, entries...)
	if err != nil {
		return err
	}
	dst := []any{}
	for _, v := range entries {
		dst = append(dst, v)
	}
	header, rows, err := table2struct.Struct2Table(dst, table2struct.WithAllowTagFields(
		"id",
		"code",
		"content",
		"result",
		"duration",
	))
	if err != nil {
		return err
	}
	allRows := [][]string{}
	allRows = append(allRows, header...)
	allRows = append(allRows, rows...)
	exporter.ExportX(out, allRows, false)
	return nil
}

func labelsFrmFile(file string) ([]hOllamaClassic.Label, error) {
	rows, err := importer.ParseFile(file)
	if err != nil {
		return nil, err
	}
	labels := []hOllamaClassic.Label{}
	for _, row := range rows {
		labels = append(labels, hOllamaClassic.Label{
			Category: row[0],
			Text:     row[1],
		})

	}
	return labels, nil
}
func textFrmFile(file string) ([]string, error) {
	rows, err := importer.ParseFile(file)
	if err != nil {
		return nil, err
	}
	res := []string{}
	for _, row := range rows {
		if len(row) < 1 {
			continue
		}
		res = append(res, row[0])
	}
	return res, nil
}
