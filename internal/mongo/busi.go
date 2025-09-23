package mongo

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/illidaris/aphroditecli/pkg/log"

	"github.com/illidaris/aphrodite/component/mongoex"
	"github.com/illidaris/aphrodite/pkg/dependency"
	"github.com/illidaris/aphrodite/pkg/group"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Demo struct {
}

func (d *Demo) ID() any {
	return ""
}
func (d *Demo) TableName() string {
	return "trade_product"
}
func (d *Demo) Database() string {
	return ""
}
func (d *Demo) ToRow() []string {
	return nil
}
func (d *Demo) ToJson() string {
	return ""
}
func (d *Demo) DbSharding(keys ...any) string {
	return "5078"
}
func (d *Demo) GetRawIndexes() []mongo.IndexModel {
	return []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "userId", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "seasonRankStr", Value: "text"},
				{Key: "name", Value: "text"},
			},
		},
	}
}

func IndexSync(ctx context.Context, dbname, conn string, concurrence int) error {
	mongoex.NewMongo("5078", dbname, conn)
	c := mongoex.GetNamedMongoClient("5078")

	mongoex.SyncDbStruct([][]any{{"5078"}}, &Demo{})
	// c.Database(dbname).Collection("trade_product").Indexes().DropOne(ctx, "userId_1_seasonRankStr_text_name_text")
	// c.Database(dbname).Collection("trade_product").Indexes().DropOne(ctx, "userId_1")
	cursor, err := c.Database(dbname).Collection("trade_product").Indexes().List(ctx)
	if err != nil {
		return err
	}
	var indexes []interface{}
	if err = cursor.All(ctx, &indexes); err != nil {
		return err
	}
	fmt.Println("Found indexes:")
	for _, idx := range indexes {
		fmt.Printf("  - %+v\n", idx)
	}
	return nil
}

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
