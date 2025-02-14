package mongo

import (
	"context"
	"encoding/json"

	"github.com/illidaris/aphrodite/component/mongoex"
	"github.com/illidaris/aphrodite/pkg/dependency"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLockRecord() LockRecord {
	return LockRecord{
		Status:     1,
		ExecAt:     0,
		Retries:    0,
		LockExpire: 0,
		NotifyAt:   0,
		NotifyEnd:  0,
		Args:       map[string]interface{}{},
	}
}

// LockRecord 推送记录表
type LockRecord struct {
	Id         string                 `json:"id" gorm:"column:id;type:varchar(32);primaryKey;comment:唯一Id" bson:"_id,omitempty"`           // id
	OpenId     string                 `json:"openId" gorm:"column:openId;type:varchar(32);comment:微信openid" bson:"openId"`                 // 微信openid
	Uid        int64                  `json:"uid" gorm:"column:uid;type:bigint;comment:uid" bson:"uid"`                                    // uid
	BizId      int64                  `json:"bizId" gorm:"column:bizId;type:bigint;comment:游戏ID" bson:"bizId"`                             // game id
	ActivityId int64                  `json:"activityId" gorm:"column:activityId;type:bigint;comment:活动ID" bson:"activityId"`              // 活动ID
	Args       map[string]interface{} `json:"args" gorm:"column:args;type:varchar(255);comment:推送参数" bson:"args"`                          // 参数
	MsgTmpId   string                 `json:"msgTmpId" gorm:"column:msgTmpId;type:varchar(64);comment:消息模板id" bson:"msgTmpId"`             // 消息模版Id
	Link       string                 `json:"link" gorm:"column:link;type:varchar(255);comment:链接地址" bson:"link"`                          // 链接地址
	Status     int32                  `json:"status" gorm:"column:status;type:int;default:1;comment:状态" bson:"status"`                     // 推送状态 0-待推送 1-已推送 2-推送失败
	CreateAt   int64                  `json:"createAt" gorm:"column:createAt;<-:create;index;autoCreateTime;comment:创建时间" bson:"createAt"` // 创建时间
	UpdateAt   int64                  `json:"updateAt" gorm:"column:updateAt;autoUpdateTime;comment:修改时间" bson:"updateAt"`                 // 修改时间
	Lock       string                 `json:"lock" gorm:"column:lock;type:varchar(36);comment:锁" bson:"lock"`                              // 锁
	TraceId    string                 `json:"traceId" gorm:"column:traceId;type:varchar(36);comment:追踪ID" bson:"traceId"`                  // 链路追踪ID
	LockExpire int64                  `json:"lockExpire" gorm:"column:lockExpire;type:bigint;comment:锁过期时间" bson:"lockExpire"`             // 锁过期时间
	ExecAt     int64                  `json:"execAt" gorm:"column:execAt;comment:执行时间" bson:"execAt,omitempty"`                            // 上一次执行时间
	ExecErr    string                 `json:"execErr" gorm:"column:execErr;comment:执行失败原因" bson:"execErr"`                                 // 上一次执行失败原因
	Retries    int64                  `json:"retries" gorm:"column:retries;type:bigint;comment:重试次数" bson:"retries"`                       // 重试次数
	NotifyAt   int64                  `json:"notifyAt" gorm:"column:notifyAt;comment:推送时间" bson:"notifyAt"`                                // 推送时间
	NotifyEnd  int64                  `json:"notifyEnd" gorm:"column:notifyEnd;comment:推送结束" bson:"notifyEnd"`                             // 推送结束【周期性活动需】
	Date       int32                  `json:"date" gorm:"column:date;type:int;comment:日期" bson:"date"`                                     // 日期, eg. 20241106
}

func (s LockRecord) TableName() string {
	return "lock_record"
}

func (s LockRecord) ID() any {
	return s.Id
}

func (s LockRecord) Database() string {
	return ""
}

func (s LockRecord) DbSharding(keys ...any) string {
	if s.BizId > 0 {
		return cast.ToString(s.BizId)
	}
	if len(keys) == 0 {
		return ""
	}
	return cast.ToString(keys[0])
}

func (s LockRecord) ToJson() string {
	bs, err := json.Marshal(&s)
	if err != nil {
		return ""
	}
	return string(bs)
}

func (s LockRecord) ToRow() []string {
	return []string{}
}

func NewLockRecordRepository() *LockRecordRepository {
	return &LockRecordRepository{}
}

type LockRecordRepository struct {
	mongoex.BaseRepository[LockRecord]
}

func (r *LockRecordRepository) DeleteActSubByID(ctx context.Context, bizId int64, id string) (int64, error) {
	var (
		opt = dependency.NewBaseOption(
			dependency.WithDbShardingKey(bizId),
		)
		filter = bson.M{"_id": id}
		count  = int64(0)
	)
	finalErr := r.BuildFrmOption(ctx, nil, opt, func(colls *mongo.Collection) error {
		res, err := colls.DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		count = res.DeletedCount
		return err
	})
	return count, finalErr

}
