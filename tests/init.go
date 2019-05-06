package tests

import (
	"encoding/json"
	"go_service/config"
	"go_service/utils"
	"log"
)

func init() {
	log.Println("Initialize config")

	// Initialize config
	utils.PanicErr(config.InitConfig("../config"))

	// Initialize DB
	utils.PanicErr(config.InitDB())
}

type M map[string]interface{}

func (t M) String() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}
