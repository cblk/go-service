package models

import (
	"encoding/json"
	"go-service/utils"
)

type M map[string]interface{}

func (t M) String() string {
	_dt, err := json.Marshal(t)
	utils.PanicErr(err)
	return string(_dt)
}
