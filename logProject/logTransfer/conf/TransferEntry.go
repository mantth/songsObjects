package conf

// TransferEntry etcd配置文件映射；
type TransferEntry struct {
	Topic string `json:"topic"`
	Index string `json:"index"`
}
