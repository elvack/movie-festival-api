package movie

import "time"

type Movie struct {
	CreatedAt   time.Time
	Description string
	Duration    int16
	Id          uint32
	Title       string
	UpdatedAt   time.Time
	WatchUrl    string
}
