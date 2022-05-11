package es

import (
	"context"
	"logTransfer/conf"
)

// es管理器；
type esManager struct {
	esTaskMap map[string]*esTask
	transList []*conf.TransferEntry
	confChan  chan []*conf.TransferEntry
}

var (
	manager *esManager
)

// watch 监听是否有新配置文件输入，获取到新的配置文件就会进行相应操作；
func (m *esManager) watch() {
	for {
		newConf := <-m.confChan
		for _, config := range newConf {
			// 如果已存在，则跳过；
			if m.isExist(config) {
				continue
			}
			// 如果不存在，则创建新的消息通道和 esTask 并将其保存到 map 中；
			newMsgChan(config.Topic)
			task := newEsTask(config.Index, config.Topic)
			m.esTaskMap[task.topic] = task
			// 启动新 esTask 给 es 发送信息；
			go task.Run(true)
		}
		// 如果在新的配置文件中未发现已存在的 esTask ，将其关闭；
		for key, task := range m.esTaskMap {
			var found bool
			for _, config := range newConf {
				if key == config.Topic {
					found = true
					break
				}
			}
			if !found {
				// 告诉该协程退出；
				task.cancel()
				// 坑：
				// 此处是否要关闭；
				closeMsgChan(key)
				delete(m.esTaskMap, key)
			}
		}
	}
}

// 判断配置文件是否已存在；
func (m *esManager) isExist(conf *conf.TransferEntry) bool {
	_, ok := m.esTaskMap[conf.Index]
	return ok
}

// Init 初始化 esTask 管理器；
func Init(entryList []*conf.TransferEntry, size int) {
	InitMsgChan(entryList, size)
	manager = &esManager{
		esTaskMap: make(map[string]*esTask, 20),
		transList: entryList,
		confChan:  make(chan []*conf.TransferEntry),
	}
	for _, config := range entryList {
		if manager.isExist(config) {
			continue
		}
		task := newEsTask(config.Index, config.Topic)
		manager.esTaskMap[task.index] = task
		go task.Run(true)
	}
	go manager.watch()
	return
}

func newEsTask(index, topic string) *esTask {
	ctx, cancelFunc := context.WithCancel(context.Background())
	task := &esTask{
		index:  index,
		topic:  topic,
		ctx:    ctx,
		cancel: cancelFunc,
	}
	return task
}

func PutNewConf(newConf []*conf.TransferEntry) {
	manager.confChan <- newConf
}
