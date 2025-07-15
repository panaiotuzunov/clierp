package main

import "time"

type entranceReceipt struct {
	id      int
	date    time.Time
	truck   string
	trailer string
	gross   int
	tare    int
	net     int
}
