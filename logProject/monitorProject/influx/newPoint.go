package influx

// NewPoint 封装消息结构体；
type NewPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
}
