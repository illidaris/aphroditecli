package watermill

import (
	"context"
	"time"

	"github.com/illidaris/watermillex/kafkaex"
)

func Publish(ctx context.Context, brokers []string, user, pwd, topic, key, value string) error {
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
	box := kafkaex.NewBoxMessage()
	box.WithOption(
		kafkaex.WithSharedGroup(), // 使用内置发布器
		kafkaex.WithTopic(topic),
		kafkaex.WithHandleTimeout(time.Minute*20),
		kafkaex.WithRetryMax(2), // 重试次数
		kafkaex.WithKey(key))    // 设置消息的key（用于分区）
	box.Value = []byte(value)
	return kafkaex.GetManager().Publish(topic, box)
}
