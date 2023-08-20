package geecache

// ByteView保存了一个不可变的字节视图。
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// 方法返回一个拷贝，防止缓存值被外部程序修改。
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String以字符串形式返回数据，必要时生成副本。
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
