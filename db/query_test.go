package db

import (
	"testing"

	"github.com/wlcsm/poll_app/common"
	"github.com/wlcsm/poll_app/db"
)

func TestCreatePoll(t *testing.T) {

	question := common.SingleChoiceQuestion{
		Header: "Hi",
	}

	questions := []common.Question{{
		Type: common.SingleChoice,
		Id:   1,
		Data: question,
	}}

	d, err := db.GetConn()
	if err != nil {
		t.Errorf("Error getting connection: %v", err)
	}

	p, err = d.CreatePoll(questions)
	if err != nil {
		t.Errorf("Error creating poll: %v", err)
	}

	r, err = d.GetPoll(p.Id)
	if err != nil {
		t.Errorf("Error getting poll: %v", err)
	}

	if r.Questions != questions {
		t.Error("Returned poll does not equal created poll")
	}
}
