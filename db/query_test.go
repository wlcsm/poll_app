package db_test

import (
	"reflect"
	"testing"

	"github.com/wlcsm/poll_app/common"
	"github.com/wlcsm/poll_app/db"
)

func TestCreatePoll(t *testing.T) {
	t.Parallel()

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

	p, err := d.CreatePoll(questions)
	if err != nil {
		t.Errorf("Error creating poll: %v", err)
	}

	r, err := d.QueryPoll(p.Id)
	if err != nil {
		t.Errorf("Error getting poll: %v", err)
	}

	if !reflect.DeepEqual(r.Questions, questions) {
		t.Errorf("Returned poll does not equal created poll\nCreated: %+v\nReturned: %+v\n", questions, r.Questions)
	}
}

func TestUpdatePoll(t *testing.T) {
	t.Parallel()

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
		t.Errorf("Error getting database connection: %v", err)
	}

	p, err := d.CreatePoll(questions)
	if err != nil {
		t.Errorf("Error creating poll: %v", err)
	}

	q := common.Question{
		Type: common.MultiChoice,
		Id:   int64(1),
		Data: question,
	}
	questions = append(questions, q)

	_, err = d.UpdatePoll(p.Id, questions)
	if err != nil {
		t.Errorf("Error updating poll %d: %v", q.Id, err)
	}

	r, err := d.QueryPoll(p.Id)
	if err != nil {
		t.Errorf("Error updating poll %d: %v", q.Id, err)
	}

	if !reflect.DeepEqual(r.Questions, questions) {
		t.Errorf("Returned poll does not equal created poll\nCreated: %+v\nReturned: %+v\n", questions, r.Questions)
	}
}

func TestDeletePoll(t *testing.T) {
	t.Parallel()

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
		t.Errorf("Error getting database connection: %v", err)
	}

	p, err := d.CreatePoll(questions)
	if err != nil {
		t.Errorf("Error creating poll: %v", err)
	}

	if err = d.DeletePoll(p.Id); err != nil {
		t.Errorf("Error deleting poll %d: %v", p.Id, err)
	}

	r, err := d.QueryPoll(p.Id)
	if err != nil {
		t.Errorf("Error querying poll %d: %v", p.Id, err)
	}
	if r != nil {
		t.Errorf("Poll wasn't deleted")
	}
}
