package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wlcsm/poll_app/common"
	"github.com/wlcsm/poll_app/db"
)

type QueryPollReq struct {
	PollId int64 `json:"poll_id"`
}

type QueryPollRes struct {
	Questions []common.Question `json:"questions"`
}

// QueryPoll
// @Router /poll [GET].
func QueryPoll(r *http.Request) common.HTTPResponse {
	var req QueryPollReq

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		return ErrResp(err)
	}

	d, err := db.GetConn()
	if err != nil {
		return ErrResp(err)
	}

	poll, err := d.QueryPoll(req.PollId)
	if err != nil {
		var code int
		if errors.Is(err, db.NotFound) {
			code = 404
		} else {
			code = 500
		}

		return common.HTTPResponse{
			Code: code,
			Err: err,
		}
	}

	return OkResp(poll)
}
