package main

import (
	conductor "github.com/hipinion/conductor/src"
)

// Setup Configs
func init() {

}

// Run server
func main() {
	conductor.Conf.Port = 9999
	conductor.Conf.ViewsDirectory = "views/*.html"
	conductor.Init()
	err := conductor.StartServer()
	if err != nil {
		panic(err)
	}
}
