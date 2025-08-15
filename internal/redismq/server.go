package redismq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "192.168.97.71:6379"

func Server() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, Password: "s8PPBYyhTkHBSyheLPdB"},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(string(ApmqMsgTypeCommon), func(ctx context.Context, t *asynq.Task) error {
		var p CommonMessage
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}
		log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
		// Email delivery code ...
		return nil
	})
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
