package handler

import (
	"net/http"

	"github.com/poll_app/common"
	"github.com/poll_app/db"
)

type SubmitPollReq struct {
	PollId  int64           `json:"poll_id"`
	Answers []common.Answer `json:"answers"`
}

// SubmitPoll Submit the answers for a poll
func SubmitPoll(w http.ResponseWriter, r *http.Request) {
	var req SubmitPollReq
	if err := ParseRequest(r, req); err != nil {
		ReturnErr(w, err, 1)
	}

	d, err := db.GetConn()
	if err != nil {
		ReturnErr(w, err, 1)
	}

	if err := d.SubmitPoll(req.PollId, req.Answers); err != nil {
		ReturnErr(w, err, 1)
	}
}
