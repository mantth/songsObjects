package conf

// KafkaConfig 配置文件结构体；
type KafkaConfig struct {
	Address string `ini:"address"`
	//Topic   string `ini:"topic"`

}

type EsConfig struct {
	Address string `ini:"address"`
	//Index   string `ini:"index"`
	Size int `ini:"size"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	Key     string `ini:"key"`
}

type Config struct {
	KafkaConf *KafkaConfig `ini:"kafka"`
	EsConf    *EsConfig    `ini:"es"`
	EtcdConf  *EtcdConfig  `ini:"etcd"`
}
