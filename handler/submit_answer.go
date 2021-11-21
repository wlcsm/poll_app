package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/wlcsm/poll_app/common"
	"github.com/wlcsm/poll_app/db"
)

type SubmitPollReq struct {
	PollId  int64           `json:"poll_id"`
	Answers []common.Answer `json:"answers"`
}

type SubmitPollRsp struct {
}

// SubmitPoll Submit the answers for a poll.
func SubmitPoll(r *http.Request) common.HTTPResponse {
	var req SubmitPollReq

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		return ErrResp(err)
	}

	resp := common.Response{Answers: req.Answers}

	d, err := db.GetConn()
	if err != nil {
		return ErrResp(err)
	}

	if err := d.SubmitPoll(req.PollId, resp); err != nil {
		fmt.Println(err)

		var code int
		if errors.Is(err, db.NotFound) {
			code = 404
		} else {
			code = 500
		}

		return common.HTTPResponse{
			Code: code,
			Err:  err,
		}
	}

	return OkResp(SubmitPollRsp{})
}
