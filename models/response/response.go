package response

// BaseResponse  基础的返回体
type BaseResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}
