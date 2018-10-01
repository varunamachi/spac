package main

import (
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"github.com/varunamachi/vaali/vsec"
)

//SprwClient - client to a sparrow server. Uses vnet.Client
type SprwClient struct {
	*vnet.Client
}

//NewClient - creates a new sprw client
func NewClient(address, versionStr string) *SprwClient {
	return &SprwClient{
		Client: vnet.NewClient(address, versionStr),
	}
}

//EntityLogin - login to entity
func (client *SprwClient) EntityLogin(creds EntityCreds) (err error) {
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
	return vlog.LogError("Spac:Client", err)
}

//Ping - ping the sparro server
func (client *SprwClient) Ping() (session *vnet.Session, err error) {
	rr := client.Get(vsec.Public, "ping")
	if rr.Err == nil {
		session = &vnet.Session{}
		err = rr.Read(session)
	} else {
		err = rr.Err
	}
	return session, vlog.LogError("Spac:Client", err)
}

//SendParamValue - send a parameter value to sparrow server
func (client *SprwClient) SendParamValue(
	paramID string,
	value string) (err error) {
	return vlog.LogError("Spac:Client", err)
}
