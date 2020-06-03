package models

import (
	"encoding/json"

	logy "github.com/sirupsen/logrus"
)

type M map[string]interface{}

func (t M) String() string {
	_dt, err := json.Marshal(t)
	if err != nil {
		logy.Error("models String error", err)
		return ""
	}

	return string(_dt)
}
