package conf

type EsConfig struct {
	Address string `ini:"address"`
}

type Config struct {
	EsConf *EsConfig `ini:"es"`
}
