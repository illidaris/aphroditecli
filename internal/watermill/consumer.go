package watermill

import (
	"context"
	"encoding/json"
	"time"

	"github.com/illidaris/aphroditecli/pkg/log"

	"github.com/illidaris/watermillex/kafkaex"
)

func Consumer(ctx context.Context, topics []string, brokers []string, user, pwd string, delay int64) error {
	kafkaex.SetName("test_client")
	kafkaex.SetGetKafkaBrokersFunc(func() []string {
		return brokers
	})
	kafkaex.SetGetKafkaUserFunc(func() string {
		return user
	})
	kafkaex.SetGetKafkaPwdFunc(func() string {
		return pwd
	})
	m := kafkaex.GetManager()
	if m == nil {
		log.Info(ctx, "kafka manager is nil")
	}
	for _, topic := range topics {
		err := m.RegisterSubscriber(context.Background(),
			topic,
			kafkaex.WithGroup("github.com/illidaris/aphroditecli"), // 主消费组，不同组可以同时消费同一条消息，一条消息只能被一个组消费一次
			kafkaex.WithTopic(topic),
			kafkaex.WithHandle(Handle(delay)))
		log.Info(context.TODO(), "github.com/illidaris/aphroditecli_subscribe_%s,%v", topic, err)
	}
	return nil
}

func Handle(delay int64) func(ctx context.Context, box *kafkaex.BoxMessage) error {
	return func(ctx context.Context, box *kafkaex.BoxMessage) error {
		bs, _ := json.Marshal(box)
		log.Info(ctx, "recv_box %s", string(bs))
		log.Info(ctx, "recv_msg %s", string(box.Value))
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
		log.Info(ctx, "exec complete %s", string(box.Value))
		return nil
	}
}
