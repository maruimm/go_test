package proto


// 需要传输的对象
type RpcObj struct {
	Id   int `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
	Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
	Ok  bool `json:"ok"`
	Id  int `json:"id"`
	Msg string `json:"msg"`
}