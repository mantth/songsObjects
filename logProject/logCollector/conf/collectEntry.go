package conf

// CollectEntry 日志收集配置；
type CollectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}
