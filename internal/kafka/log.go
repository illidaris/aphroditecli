package kafka

import (
	"context"
	"fmt"
)

type ConsoleLogger struct {
}

func (l *ConsoleLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}
func (l *ConsoleLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}
func (l *ConsoleLogger) Println(v ...interface{}) {
	fmt.Println(v...)
}
func (l *ConsoleLogger) Debug(_ context.Context, msg string, args ...interface{}) {
	fmt.Printf("[DEBUG]"+msg+"\n", args...)
}
func (l *ConsoleLogger) Info(_ context.Context, msg string, args ...interface{}) {
	fmt.Printf("[INFO]"+msg+"\n", args...)
}
func (l *ConsoleLogger) Warn(_ context.Context, msg string, args ...interface{}) {
	fmt.Printf("[WARN]"+msg+"\n", args...)
}
func (l *ConsoleLogger) Error(_ context.Context, msg string, args ...interface{}) {
	fmt.Printf("[ERROR]"+msg+"\n", args...)
}
