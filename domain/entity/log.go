package entity

type Log struct {
	CreatedAt uint64 `json:"createdAt"`
	Error     string `json:"error"`
}
