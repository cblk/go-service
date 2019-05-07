package tests

import (
	"encoding/json"
	"log"

	"go_service/config"
)

func init() {
	log.Println("Initialize config")

	// Initialize config
	err := config.InitConfig("../config")
	if err != nil {
		log.Println("init config error:",err.Error())
		return
	}

	// Initialize DB
	err = config.InitDB()
	if err != nil {
		log.Println("init DB error:",err.Error())
		return
	}
}

type M map[string]interface{}

func (t M) String() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}
