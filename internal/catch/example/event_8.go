package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go_service/library/logy"
)

const (
	CONST_VisitThird_MaxTime = 30 // s
)

var (
	ERR_NetWork        = errors.New("network error")
	ERR_CLOSED_NetWork = errors.New("io: read/write on closed pipe")
)

// mock 不断重试的机制
func VisitThirdEvent(params ...string) error {
	if len(params) == 0 {
		return ERR_Params
	}

	trySleepTime := 1
tryVisitThird:
	resp, err := http.Get("http://****:8080/api/task")
	if err != nil {
		logy.Error("VisitThirdEvent", err)

		// try
		waitTime := time.Duration(trySleepTime) * time.Second
		time.Sleep(waitTime)

		waitTime = waitTime * 2
		if waitTime >= CONST_VisitThird_MaxTime {
			waitTime = CONST_VisitThird_MaxTime
		}

		goto tryVisitThird
	}

	content := resp.Body
	// do something
	log.Println(content)

	return nil
}
