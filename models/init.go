package models

import (
	"encoding/json"

	"go_service/library/logy"
)

type M map[string]interface{}

func (t M) String() string {
	_dt, err := json.Marshal(t)
	if err != nil {
		logy.ErrorW("models String error", err).Error()
		return ""
	}

	return string(_dt)
}
