package handler

import (
	"net/http"

	"github.com/poll_app/common"
	"github.com/poll_app/db"
)

type CreatePollReq struct {
	Questions []common.Question `json:"questions"`
}

type CreatePollRsp struct {
	PollId int64 `json:"poll_id"`
}

// CreatePoll Creates a poll
// @Success 200 {object} CreatePollRsp
// @Router /poll/create [POST]
func CreatePoll(w http.ResponseWriter, r *http.Request) {
	var req CreatePollReq
	if err := ParseRequest(r, req); err != nil {
		ReturnErr(w, err, 1)
	}

	d, err := db.GetConn()
	if err != nil {
		ReturnErr(w, err, 1)
	}

	poll, err := d.CreatePoll(
		[]common.Question{{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		},
		})
	if err != nil {
		ReturnErr(w, err, 1)
	}

	ReturnResp(w, CreatePollRsp{poll.Id}, nil, 1)
}
