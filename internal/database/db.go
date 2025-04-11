package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// DbQuery 执行sql语句
func DbQuery(ctx context.Context, sqlstr string, args []any, opts ...Option) ([]map[string]any, error) {
	argStrs := []string{}
	for _, arg := range args {
		argStrs = append(argStrs, cast.ToString(arg))
	}
	println(fmt.Sprintf("执行sql语句耗时：%v, [%s]", sqlstr, strings.Join(argStrs, ",")))
	result := []map[string]any{}
	o := NewOptions(opts...)
	db, err := sql.Open(o.Driver, o.DSN)
	if err != nil {
		return result, err
	}
	defer db.Close()
	now := time.Now()
	result, err = RetrieveMap(ctx, db, sqlstr, args...)
	println(fmt.Sprintf("执行sql语句耗时：%vms", time.Since(now).Milliseconds()))
	if err != nil {
		return result, err
	}
	return result, nil
}

// RetrieveMap SQL查询结果输出为Map
func RetrieveMap(ctx context.Context, db *sql.DB, sSql string, args ...any) ([]map[string]any, error) {
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
	mData := make([]map[string]interface{}, 0)
	// 一条数据的各列的值（需要指定长度为列的个数，以便获取地址）
	values := make([]interface{}, count)
	// 一条数据的各列的值的地址
	valPointers := make([]interface{}, count)
	for rows.Next() {
		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}
		// 获取各列的值，放到对应的地址中
		rows.Scan(valPointers...)
		// 一条数据的Map (列名和值的键值对)
		entry := make(map[string]interface{})
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
			entry[col] = v
		}
		mData = append(mData, entry)
	}
	return mData, nil
}
