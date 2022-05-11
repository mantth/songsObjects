package tails

import (
	"context"
	"log"
	"logProject/conf"
)

// tail管理器；
type tailManager struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []*conf.CollectEntry
	confChan         chan []*conf.CollectEntry
}

// 监听配置文件变化，对tailTask做相应更改；
func (m *tailManager) watch() {
	for {
		newConf := <-m.confChan
		for _, logConfig := range newConf {
			// 如果是已存在的配置，则直接跳过；
			if m.isExist(logConfig) {
				continue
			}
			// 存在新配置项，启动新的tailTask；
			task := newTailTask(logConfig.Path, logConfig.Topic)
			err := task.Init()
			if err != nil {
				log.Printf("create tail failed: %s", err)
				continue
			}
			// 将新的tailTask保存到列表中；
			m.tailTaskMap[task.path] = task
			go task.Run()
		}
		// 如果当前配置中存在，但新配置中不存在，则kill;
		for key, task := range m.tailTaskMap {
			var found bool
			for _, config := range newConf {
				if key == config.Path {
					found = true
					break
				}
			}
			if !found {
				//err := task.tailObj.Stop()
				//if err != nil {
				//	fmt.Println(err)
				//	return
				//}
				// 不存在，则停止该goroutine;
				task.cancel()
				//m.tailTaskMap[task.path] =
				delete(m.tailTaskMap, key)
			}
		}
	}
}

// 判断配置文件是否已存在；
func (m *tailManager) isExist(conf *conf.CollectEntry) bool {
	_, ok := m.tailTaskMap[conf.Path]
	return ok
}

var (
	manager *tailManager
)

func Init(entryList []*conf.CollectEntry) (err error) {
	// 初始化全局管理器；
	manager = &tailManager{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: entryList,
		confChan:         make(chan []*conf.CollectEntry),
	}
	// 对每个日志收集想启用一个协程
	for _, logConfig := range entryList {
		task := newTailTask(logConfig.Path, logConfig.Topic)
		err = task.Init()
		if err != nil {
			log.Printf("create tail failed: %s", err)
			continue
		}
		manager.tailTaskMap[task.path] = task
		go task.Run() // 把task存储起来；

	}
	// 起单独协程，监听配置文件变化；
	go manager.watch()
	return
}

// 创建新的tailTask；
func newTailTask(path, topic string) *tailTask {
	ctx, cancelFunc := context.WithCancel(context.Background())
	task := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancelFunc,
	}
	return task
}

// PutNewConf 开放给etcd使用；
func PutNewConf(newConf []*conf.CollectEntry) {
	manager.confChan <- newConf
}
