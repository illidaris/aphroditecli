package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/illidaris/aphrodite/component/kafkaex"
)

func Consumer(ctx context.Context, topics []string, opts ...kafkaex.OptionsFunc) error {
	kafkaex.SetLogger(&ConsoleLogger{})
	m, err := kafkaex.InitDefaultManager(opts...)
	if err != nil {
		return err
	}
	err = m.NewConsumer(uuid.NewString(), "aphroditecli", func(ctx context.Context, m *kafkaex.Message) (kafkaex.ReceiptStatus, error) {
		msg := fmt.Sprintf("%s[消费:%s]：消费者(%s|%s),主题(%s_%d_%d),消息头(%s),%s=%s \n",
			time.Now().Format("2006-01-02 15:04:05"),
			time.Unix(m.Ts, 0).Format("2006-01-02 15:04:05"),
			m.Id,
			m.ConsumerId,
			m.Topic,
			m.Partition,
			m.Offset,
			m.Headers,
			string(m.Key),
			string(m.Value))
		println(msg)
		return kafkaex.ReceiptSuccess, nil
	}, topics...)
	if err != nil {
		return err
	}
	m.ConsumersGo()
	return nil
}
