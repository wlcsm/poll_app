package db

import (
	"fmt"

	"github.com/wlcsm/poll_app/common"
)

type DB interface {
	QueryPoll(pollId int64) (*common.Poll, error)
	CreatePoll(q []common.Question) (*common.Poll, error)
	UpdatePoll(pollId int64, q []common.Question) (*common.Poll, error)
	DeletePoll(pollId int64) error
	SubmitPoll(pollId int64, resp common.Response) error
}

// Mock database for now.
var d = mock{
	Polls: make(map[int64]common.Poll),
}

type mock struct {
	Polls     map[int64]common.Poll
	Responses map[int64][]common.Response
}

func GetConn() (DB, error) {
	return &d, nil
}

// QueryPoll
// Query the poll.
func (d *mock) QueryPoll(pollId int64) (*common.Poll, error) {
	res, ok := d.Polls[pollId]
	if !ok {
		return nil, nil
	}

	return &res, nil
}

// CreatePoll
// Create the poll. Return an error if it already exists.
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

// UpdatePoll
// Update the poll. Returns an error if it doesn't exist.
func (d *mock) UpdatePoll(pollId int64, q []common.Question) (*common.Poll, error) {
	p, ok := d.Polls[pollId]
	if !ok {
		return nil, NotFound
	}

	p.Questions = q
	d.Polls[pollId] = p

	return &p, nil
}

// DeletePoll
// Delete the poll matching the PollQuery from the database.
func (d *mock) DeletePoll(pollId int64) error {
	delete(d.Polls, pollId)
	return nil
}

// SubmitPoll
// Submit a response to the poll.
func (d *mock) SubmitPoll(pollId int64, resp common.Response) error {
	_, ok := d.Polls[pollId]
	if !ok {
		return fmt.Errorf("submitting a response to a poll that doesn't exists: %w", NotFound)
	}

	res, ok := d.Responses[pollId]
	if !ok {
		res = []common.Response{}
	}

	res = append(res, resp)
	d.Responses[pollId] = res

	return nil
}

const MAX_DB = 10000

// Get an available ID.
func (d *mock) GetAnId(q []common.Question) *int64 {
	for i := int64(0); i < MAX_DB; i++ {
		if _, ok := d.Polls[i]; !ok {
			return &i
		}
	}

	return nil
}
