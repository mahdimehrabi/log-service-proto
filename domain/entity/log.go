package entity

import "time"

type Log struct {
	CreatedAt uint64 `json:"createdAt"`
	Error     string `json:"error"`
}

func NewLog(err string) *Log {
	return &Log{
		Error:     err,
		CreatedAt: uint64(time.Now().Unix()),
	}
}
