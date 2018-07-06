package main

import (
	"github.com/varunamachi/vaali/vapp"
)

//GetModule - gets module corresponding to spac
func GetModule() *vapp.Module {
	return &vapp.Module{
		Name:     "spac",
		Commands: GetCommands(),
	}
}
