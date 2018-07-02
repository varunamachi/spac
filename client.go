package main

import (
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
)

//EntityLogin - login to entity
func EntityLogin(creds EntityCreds) (err error) {
	return vlog.LogError("Spc:Client", err)
}

//Ping - ping the sparrow server
func Ping() (info vnet.M, err error) {
	return info, vlog.LogError("Spc:Client", err)
}

//SendParamValue - send a parameter value to sparrow server
func SendParamValue(paramID string, value string) (err error) {
	return vlog.LogError("Spc:Client", err)
}
