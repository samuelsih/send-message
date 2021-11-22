package main

import (
	"github.com/samuelssh/send-message/util"
	"os"
)

func main() {
	util.LoadENV()
	util.CreateClientServer()
	util.SendMessage(os.Getenv("MESSAGE"))
}


