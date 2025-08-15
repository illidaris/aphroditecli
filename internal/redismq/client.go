package redismq

import (
	"log"

	"github.com/hibiken/asynq"
)

func Client() {
	client := asynq.NewClient(
		asynq.RedisClientOpt{Addr: redisAddr, Password: "s8PPBYyhTkHBSyheLPdB"},
	)
	defer client.Close()

	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately.
	//            Use (*Client).Enqueue method.
	// ------------------------------------------------------
	m := &CommonMessage{}
	m.UserID = 41
	m.TemplateID = "xxxx2"

	task, err := NewPayloadTask(m)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ------------------------------------------------------------
	// Example 2: Schedule task to be processed in the future.
	//            Use ProcessIn or ProcessAt option.
	// ------------------------------------------------------------

	// info, err = client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	// if err != nil {
	// 	log.Fatalf("could not schedule task: %v", err)
	// }
	// log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

}
