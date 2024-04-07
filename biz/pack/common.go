package pack

import "videoweb/biz/model/videoweb"

func BuildErrResp(code int, err error) videoweb.ErrResp {
	return videoweb.ErrResp{
		Base: &videoweb.BaseResp{
			Code: int64(code),
			Msg:  err.Error(),
		},
	}
}
