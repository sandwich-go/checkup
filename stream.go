package internalcmd

import "time"

const MetaKeyStreamID = "sandwich_stream_id"

type IStream interface {
	GetStreamAddr() string
	NewStreamToken() string
	GetCache() IStreamCache
}

type IStreamCache interface {
	// Get 返回key对应的元素
	Get(key interface{}) (interface{}, bool, error)
	// Set 设定元素，使用携带的ttl
	// 若ttl = 0，则使用默认的DefaultTTL，否则使用ttl
	Set(key interface{}, value interface{}, ttl time.Duration)
	// Del 删除元素
	Del(key interface{}) (bool, error)
}

type StreamArgs struct {
	ID    string
	Addr  string
	Token string
	Meta  map[string]string
}
