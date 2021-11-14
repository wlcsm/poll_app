package handler

import (
	"net/http"

	"github.com/poll_app/common"
	"github.com/poll_app/db"
)

type GetPollReq struct {
	PollId int64 `json:"poll_id"`
}

type GetPollRes struct {
	Questions []common.Question `json:"questions"`
}

// GetPoll
// @Router /poll [GET]
func GetPoll(w http.ResponseWriter, r *http.Request) {
	d, err := db.GetConn()
	if err != nil {
		ReturnErr(w, err, 1)
	}

	poll, err := d.QueryPoll(0)
	if err != nil {
		ReturnErr(w, err, 1)
	}

	ReturnResp(w, poll, nil, 0)
}
