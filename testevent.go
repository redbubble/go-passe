package main

import "time"

type TestEvent struct {
	Time    time.Time
	Action  string
	Elapsed float64
	Output  string
	TestID
}
