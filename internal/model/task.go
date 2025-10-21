package model

type task struct {
	id          int64
	description string
	status      string
	createdAt   timestamp
	updatedAt   timestamp
}
