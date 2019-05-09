package main

var _TenantId string

// 之前
func SetTenantId_(val string) error {
	_TenantId = val

	return nil
}

// 改进后
func SetTenantId(val string) {
	_TenantId = val
}