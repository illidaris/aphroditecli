package mongo

import (
	"aphroditecli/pkg/log"
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/illidaris/aphrodite/component/mongoex"
	"github.com/illidaris/aphrodite/pkg/dependency"
	"github.com/illidaris/aphrodite/pkg/group"
	"go.mongodb.org/mongo-driver/bson"
)

func Exec(ctx context.Context, dbname, conn string, concurrence int) error {
	mongoex.NewMongo("5078", dbname, conn)

	ids := []int64{}
	for i := 0; i < concurrence; i++ {
		ids = append(ids, int64(i))
	}
	repo := NewLockRecordRepository()
	_, _ = repo.BaseCreate(ctx, []*LockRecord{
		{
			Uid: 777,
		},
		{
			Uid: 88,
		},
	}, dependency.WithDbShardingKey(5078))

	count := int32(0)
	countPtr := &count

	_, errMap := group.GroupBaseFunc(func(v ...int64) (int64, error) {

		// uid := v[0]
		filter := bson.D{}
		filter = append(filter, bson.E{Key: "uid", Value: 1})
		opts := []dependency.BaseOptionFunc{
			dependency.WithDbShardingKey(5078),
			dependency.WithConds(filter),
		}
		records, err := repo.BaseQuery(ctx, opts...)
		if len(records) > 0 {
			println(records[0].Id)
		}
		time.Sleep(time.Millisecond * 500)
		atomic.AddInt32(countPtr, 1)
		return int64(len(records)), err
	}, 2, ids...)
	log.Info(ctx, "%v %v", count, errMap)
	return nil
}

func ExecWithTrans(ctx context.Context, dbname, conn string, concurrence int) error {
	mongoex.NewMongo("5078", dbname, conn)

	uow := mongoex.NewMongoUnitOfWork("5078")
	repo := NewLockRecordRepository()
	// uid := v[0]
	opts := []dependency.BaseOptionFunc{
		dependency.WithDbShardingKey(5078),
	}

	actions := []dependency.DbAction{
		func(subCtx context.Context) error {
			affect, err := repo.BaseCreate(subCtx, []*LockRecord{
				{
					Uid: 777,
				},
			}, opts...)
			println("事务子任务执行：", affect)
			return err
		},
		func(subCtx context.Context) error {
			affect, err := repo.BaseCreate(subCtx, []*LockRecord{
				{
					Uid: 88,
				},
			}, opts...)
			println("事务子任务执行：", affect)
			return err
		},
		func(subCtx context.Context) error {
			panic("失败啦~~~")
			return errors.New("退出")
		},
	}
	err := uow.Execute(ctx, actions...)
	log.Info(ctx, "%v", err)
	return nil
}
