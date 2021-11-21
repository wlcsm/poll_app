package handler

import "github.com/wlcsm/poll_app/common"

func ErrResp(err error) common.HTTPResponse {
	return common.HTTPResponse{
		Code: 500,
		Err: err,
	}
}

func OkResp(data interface{}) common.HTTPResponse {
	return common.HTTPResponse{
		Code: 200,
		Data: data,
	}
}
