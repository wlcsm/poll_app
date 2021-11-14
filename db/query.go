package db

import (
	"github.com/poll_app/common"
)

// Mock database for now

type DB interface {
	QueryPoll(pollId int64) (*common.Poll, error)
	CreatePoll(q []common.Question) (*common.Poll, error)
	DeletePoll(pollId int64) (*common.Poll, error)
	SubmitPoll(pollId int64, ans []common.Answer) error
}

var d mock

type mock struct {
	Polls map[int64]common.Poll
}

func GetConn() (DB, error) {
	return &d, nil
}

// QueryPoll
// Query the poll
func (d *mock) QueryPoll(pollId int64) (*common.Poll, error) {
	res, ok := d.Polls[pollId]
	if !ok {
		return nil, nil
	}

	return &res, nil
}

// UpsertPoll
// Insert the poll if it doesn't exist, update if it does
func (d *mock) CreatePoll(q []common.Question) (*common.Poll, error) {
	id := d.GetAnId(q)
	if id == nil {
		return nil, DBFull
	}

	d.Polls[*id] = common.Poll{
		Id:        *id,
		Questions: q,
	}

	p := d.Polls[*id]
	return &p, nil
}

// Get an available ID
func (d *mock) GetAnId(q []common.Question) *int64 {
	for i := int64(0); i < 10000; i++ {
		if _, ok := d.Polls[i]; !ok {
			return &i
		}
	}

	return nil
}

// DeletePoll
// Delete the poll matching the PollQuery from the database
func (d *mock) DeletePoll(pollId int64) (*common.Poll, error) {
	return &common.Poll{
		Questions: []common.Question{{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		}}}, nil
}

// SubmitPoll
// Submit a response to the poll
func (d *mock) SubmitPoll(pollId int64, ans []common.Answer) error {
	return nil
}
