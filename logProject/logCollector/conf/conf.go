package conf

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	LogPath string `ini:"collect_key"`
}

type Config struct {
	KafakConfig   *KafkaConfig   `ini:"kafka"`
	CollectConfig *CollectConfig `ini:"collect"`
	EtcdConfig    *EtcdConfig    `ini:"etcd"`
}
