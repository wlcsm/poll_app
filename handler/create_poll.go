package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wlcsm/poll_app/common"
	"github.com/wlcsm/poll_app/db"
)

type CreatePollReq struct {
	Questions []common.Question `json:"questions"`
}

type CreatePollRsp struct {
	PollId int64 `json:"poll_id"`
}

// CreatePoll Creates a poll
// @Success 200 {object} CreatePollRsp
// @Router /poll/create [POST].
func CreatePoll(r *http.Request) common.HTTPResponse {
	var req CreatePollReq

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		return ErrResp(err)
	}

	d, err := db.GetConn()
	if err != nil {
		return ErrResp(err)
	}

	q := []common.Question{
		{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		},
	}

	poll, err := d.CreatePoll(q)
	if err != nil {
		return ErrResp(err)
	}

	data := CreatePollRsp{poll.Id}
	return OkResp(data)
}
