package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/illidaris/aphrodite/component/gormex"
	"github.com/illidaris/aphrodite/pkg/dependency"
)

const (
	DB_ID string = "apCli"
)

func DbExec(dsn string, trans int32, delay time.Duration, sqls ...string) error {
	db, err := gormex.NewMySqlClient(dsn, gormex.NewLogger())
	if err != nil {
		return err
	}
	gormex.MySqlComponent.NewWriter(DB_ID, db)
	ctx := context.Background()
	results := []string{}
	actions := []dependency.DbAction{}
	for _, v := range sqls {
		sqlstr := v
		actions = append(actions, func(subCtx context.Context) error {
			now := time.Now()
			client := gormex.CoreFrmCtx(subCtx, DB_ID)
			result := client.Exec(sqlstr)
			results = append(results, fmt.Sprintf("%v\n执行结果[%.2f]：\nRow:%v,Err:%v", sqlstr, time.Since(now).Seconds(), result.RowsAffected, result.Error))
			if delay.Seconds() > 1 {
				time.Sleep(delay)
			}
			return nil
		})
	}
	return WitnTransOrNot(ctx, trans, actions...)
}

func WitnTransOrNot(ctx context.Context, trans int32, actions ...dependency.DbAction) error {
	if trans < 0 {
		for _, act := range actions {
			act(ctx)
		}
		return nil
	}
	opts := []gormex.UOWOptionFunc{}
	if trans > 0 {
		opts = append(opts, gormex.WithIsolationLevel(sql.IsolationLevel(trans)))
	}
	uow := gormex.NewUnitOfWork(DB_ID, opts...)
	return uow.Execute(ctx, actions...)
}
