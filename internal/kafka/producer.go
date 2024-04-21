package kafka

import (
	"context"

	"github.com/illidaris/aphrodite/component/kafkaex"
)

func Publish(ctx context.Context, topic string, key string, message string, opts ...kafkaex.OptionsFunc) error {
	kafkaex.SetLogger(&ConsoleLogger{})
	m, err := kafkaex.InitDefaultManager(opts...)
	if err != nil {
		return err
	}
	err = m.Publish(ctx, topic, "msg", []byte("xxxx"))
	if err != nil {
		return err
	}
	return nil
}
