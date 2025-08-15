package redismq

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

// https://github.com/hibiken/asynq

type MessageType string

// A list of task types.
const (
	ApmqMsgTypeCommon MessageType = "apmq:common"
)

type IMessage interface {
	MessageType() MessageType
}

type CommonMessage struct {
	UserID     int
	TemplateID string
}

func (m CommonMessage) MessageType() MessageType {
	return ApmqMsgTypeCommon
}

func NewPayloadTask(m IMessage) (*asynq.Task, error) {
	payload, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		string(m.MessageType()),
		payload,
		asynq.MaxRetry(5),
		asynq.Timeout(20*time.Minute),
	), nil
}
