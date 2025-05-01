package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/illidaris/aphrodite/pkg/group"
	"github.com/illidaris/aphrodite/pkg/structure"
	"github.com/illidaris/aphroditecli/pkg/exporter"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
)

func DbQueryExport(ctx context.Context, sqls []string, out string, pretty bool, opts ...Option) error {
	o := NewOptions(opts...)
	db, err := sql.Open(o.Driver, o.DSN)
	if err != nil {
		return err
	}
	defer db.Close()
	bar := progressbar.Default(int64(len(sqls)))
	_, _ = group.GroupBaseFunc(func(vs ...string) (int64, error) {
		for _, v := range vs {
			defer bar.Add(1)
			data, err := DbQuery(ctx, v, nil, db)
			if err != nil {
				println(err.Error())
			}
			// TODO: 导出文件名称
			exporter.Export(uuid.NewString(), out, data, pretty)
		}
		return 1, nil
	}, 1, sqls...)
	return err
}

// DbQuery 执行sql语句
func DbQuery(ctx context.Context, sqlstr string, args []any, db *sql.DB) ([]*structure.KVs[string, any], error) {
	argStrs := []string{}
	for _, arg := range args {
		argStrs = append(argStrs, cast.ToString(arg))
	}
	println(fmt.Sprintf("执行sql语句耗时：%v, [%s]", sqlstr, strings.Join(argStrs, ",")))
	result := []*structure.KVs[string, any]{}
	now := time.Now()
	result, err := RetrieveMap(ctx, db, sqlstr, args...)
	println(fmt.Sprintf("执行sql语句耗时：%vms", time.Since(now).Milliseconds()))
	if err != nil {
		return result, err
	}
	return result, nil
}

// RetrieveMap SQL查询结果输出为Map
func RetrieveMap(ctx context.Context, db *sql.DB, sSql string, args ...any) ([]*structure.KVs[string, any], error) {
	stmt, err := db.PrepareContext(ctx, sSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()
	// 查询
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	// 数据列
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 列的个数
	count := len(columns)
	// 返回值 Map切片
	result := []*structure.KVs[string, any]{}
	// 一条数据的各列的值（需要指定长度为列的个数，以便获取地址）
	values := make([]interface{}, count)
	// 一条数据的各列的值的地址
	valPointers := make([]interface{}, count)
	for rows.Next() {
		kvs := structure.NewKVs[string, any]()
		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}
		// 获取各列的值，放到对应的地址中
		rows.Scan(valPointers...)
		// Map 赋值
		for i, col := range columns {
			var v interface{}
			// 值复制给val(所以Scan时指定的地址可重复使用)
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				// 字符切片转为字符串
				v = string(b)
			} else {
				v = val
			}
			kvs.Set(col, v)
		}
		result = append(result, kvs)
	}
	return result, nil
}
