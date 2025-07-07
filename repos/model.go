package repos

import "time"

type Histories struct {
	Id       int64     `db:"id"`
	Status   string    `db:"status"`
	Detected time.Time `db:"detected"`
}
