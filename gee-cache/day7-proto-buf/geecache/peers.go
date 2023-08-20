package geecache

// PeerPicker是必须实现定位的接口
// 拥有特定的key。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// 方法用于从对应 group 查找缓存值
	Get(group string, key string) ([]byte, error)
}
