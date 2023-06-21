package errorx

// 110代表用户系统  1101代表错误码
var ParamsError = New(1101101, "参数错误")

type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func (e *BizError) Error() string {
	return e.Msg
}

func (e *BizError) Data() *ErrorResponse {
	return &ErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
