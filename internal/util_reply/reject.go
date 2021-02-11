package replyU

import (
	xerr "github.com/goclub/error"
	xjson "github.com/goclub/json"
)

type Response struct {
	Type ResponseType
	Msg string
}
type ResponseType string
func(ResponseType) Enum() (e struct{
	Pass ResponseType
	Fail ResponseType
}) {
	e.Pass = "pass"
	e.Fail = "fail"
	return
}
func RejectMessage(msg string, shouldRecord bool) error {
	resp := Response{
		Type: ResponseType("").Enum().Fail,
		Msg: msg,
	}
	data, err := xjson.Marshal(resp) ; if err != nil {
		return err
	}
	return xerr.NewReject(data, shouldRecord)
}