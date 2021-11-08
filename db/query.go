package db

import (
	"github.com/poll_app/common"
)

// Mock database for now

type PollQuery struct {
	PollId int64 `json:"poll_id"`
}

type DB interface {
	QueryPoll(q PollQuery) (common.Poll, error)
	UpdatePoll(q common.Poll) (common.Poll, error)
	DeletePoll(q PollQuery) (common.Poll, error)
	SubmitPoll(pollId int64, ans []common.Answer) error
}

func GetConn() (DB, error) {
	return &db{}, nil
}

type db struct {
}

// QueryPoll
// Query the poll
func (d *db) QueryPoll(q PollQuery) (common.Poll, error) {
	return common.Poll{
		Questions: []common.Question{{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		}}}, nil
}

// UpdatePoll
// Update the poll
func (d *db) UpdatePoll(q common.Poll) (common.Poll, error) {
	return common.Poll{
		Questions: []common.Question{{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		}}}, nil
}

// DeletePoll
// Delete the poll matching the PollQuery from the database
func (d *db) DeletePoll(q PollQuery) (common.Poll, error) {
	return common.Poll{
		Questions: []common.Question{{
			Type: common.YesOrNo,
			Id:   1,
			Data: "his",
		}}}, nil
}

// SubmitPoll
// Submit a response to the poll
func (d *db) SubmitPoll(pollId int64, ans []common.Answer) error {
	return nil
}
