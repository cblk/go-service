package main

import "errors"

// 之前
func CheckHostType(val string) error {
	switch val {
	case "machine":
		return nil
	case "metal":
		return nil
	}
	return errors.New("CheckHostType error:" + val)
}

// 改进后
func IsValidHostType(val string) bool {
	return val == "virtual_machine" || val == "bare_metal"
}

// 没有错误返回
func DoEvent() {
	// do something and no error
	return
}
