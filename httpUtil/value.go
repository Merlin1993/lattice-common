package httpUtil

// JSONRequest 请求格式结构
type JSONRequest struct {
	JsonRpc string        `json:"jsonRpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

// JSONReply 获得结果结构
type JSONReply struct {
	JsonRpc string      `json:"jsonRpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
	Error   JsonErr     `json:"error"`
}

// JsonErr 错误结构
type JsonErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
