package dto

type RespOpt struct {
	Success bool
	Message string
}

func NewApiResp[T any](data T, opt ...RespOpt) GlobalHandlerResp[T] {
	msg := "success"
	success := true

	for _, o := range opt {
		if o.Success {
			success = o.Success
		}
		if o.Message != "" {
			msg = o.Message
		}
	}

	return GlobalHandlerResp[T]{
		Success: success,
		Message: msg,
		Data:    data,
	}
}
