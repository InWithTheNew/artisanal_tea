package service

import "time"

type config struct {
	serviceName    string
	environments   []string
	commandHistory []command
}

type command struct {
	command string
	time    time.Time
}
