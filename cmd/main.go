package main

import (
	"github.com/5107293001/contact-api/internals/server"
)

func main() {
	s := server.GetServer()

	s.Run()

}
