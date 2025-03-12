package common

import "m3u82mp4/consts/errcode"

type Respone struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *Respone) OK(data any) *Respone {
	r.Code = 200
	r.Data = data
	return r
}

func (r *Respone) OKM(data any, msg string) *Respone {
	r.Code = 200
	r.Data = data
	r.Msg = msg
	return r
}

func (r *Respone) OK2() *Respone {
	r.Code = 200
	return r
}

func (r *Respone) Set(e errcode.ErrorCode) *Respone {
	r.Code = e.Code
	r.Msg = e.Msg
	return r
}

func (r *Respone) Error(message string) *Respone {
	r.Code = 1002
	r.Msg = message
	return r
}
