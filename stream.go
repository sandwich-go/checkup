package internalcmd

import (
	"github.com/sandwich-go/boost/z"
	"github.com/sandwich-go/logbus"
	"io"
	"net"
	"time"
)

const MetaKeyStreamID = "sandwich_stream_id"
const streamTokenLen = 16

type IStream interface {
	GetStreamAddr() string
	GetStreamHandler() StreamHandler
	NewStreamToken() string
	GetCache() IStreamCache
}

type IStreamCache interface {
	// Get 返回key对应的元素
	Get(key interface{}) (interface{}, bool, error)
	// Set 设定元素，使用携带的ttl
	// 若ttl = 0，则使用默认的DefaultTTL，否则使用ttl
	Set(key interface{}, value interface{}, ttl time.Duration) error
	// Del 删除元素
	Del(key interface{}) (bool, error)
}

type StreamArgs struct {
	ID    string
	Addr  string
	Token string
	Meta  map[string]string
}

func OnStream(conn net.Conn) {
	if tc, ok := conn.(*net.TCPConn); ok {
		tc.SetKeepAlive(true)
		tc.SetKeepAlivePeriod(3 * time.Minute)
		tc.SetLinger(10)
	}

	token := make([]byte, streamTokenLen)
	_, err := io.ReadFull(conn, token)
	if err != nil {
		_ = conn.Close()
		logbus.Error("stream: failed to read token", logbus.String("addr", conn.RemoteAddr().String()))
		return
	}
	tokenStr := z.BytesToString(token)
	cache := GetOptions().GetIStream().GetCache()
	streamArg, ok, _ := cache.Get(tokenStr)
	if !ok {
		_ = conn.Close()
		logbus.Error("stream: token not valid", logbus.String("token", tokenStr))
		return
	}
	_, _ = cache.Del(tokenStr)
	go GetOptions().GetIStream().GetStreamHandler()(conn, streamArg.(*StreamArgs))
}
