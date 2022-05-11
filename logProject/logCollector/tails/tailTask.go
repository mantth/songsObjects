package tails

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"log"
	"logProject/kafka"
	"strings"
	"time"
)

type tailTask struct {
	path    string
	topic   string
	tailObj *tail.Tail
	// 用于控制tailTask的运行
	ctx    context.Context
	cancel context.CancelFunc
}

// Init tailTask的初始化；
func (task *tailTask) Init() (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	task.tailObj, err = tail.TailFile(task.path, config)
	return
}

func (task *tailTask) Run() {
	for {
		select {
		case <-task.ctx.Done(): // 只要 manager 调用 ctx.cancel() 就会停止该goroutine;
			log.Printf("task %s has stopped...", task.path)
			return
		case line, ok := <-task.tailObj.Lines:
			if !ok {
				log.Printf("tail file close reopen, filename: %s", task.path)
				time.Sleep(time.Second)
				// 不ok就直接跳过
				continue
			}
			// 如果是空行，就跳过；
			if len(strings.Trim(line.Text, "\r")) == 0 {
				continue
			}
			msg := &sarama.ProducerMessage{}
			msg.Topic = task.topic
			msg.Value = sarama.StringEncoder(line.Text)
			kafka.SendMsgToChan(msg)
		}
	}
}
