package main

import (
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

var client *vnet.Client

//EntityLogin - login to entity
func EntityLogin(creds EntityCreds) (err error) {
	rr := client.Post(creds, vsec.Public, "entity/auth")
	loginResult := struct {
		Token string     `json:"token"`
		User  *vsec.User `json:"user"`
	}{}
	err = rr.Read(&loginResult)
	if err == nil {
		client.Token = loginResult.Token
		client.User = loginResult.User
	}
	return err
	return vlog.LogError("Spac:Client", err)
}

//Ping - ping the sparro server
func Ping() (info vnet.M, err error) (err error, session *vnet.Session){
	rr := client.Get(vsec.Public, "ping")
	if rr.Err == nil {
		session := &vnet.Session{}
		err == rr.Read(session)
	} else {
		err = rr.Err
	}
	return info, vlog.LogError("Spac:Client", err)
}

//SendParamValue - send a parameter value to sparrow server
func SendParamValue(paramID string, value string) (err error) {
	return vlog.LogError("Spac:Client", err)
}
