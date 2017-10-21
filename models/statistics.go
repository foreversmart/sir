package models

type Statistics struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`

	*TaskState
}
