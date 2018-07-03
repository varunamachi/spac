package main

//EntityCreds - creds used to authenticate an entity
type EntityCreds struct {
	EntityID string `json:"entityID"`
	Owner    string `json:"owner"`
	Secret   string `json:"secret"`
}
