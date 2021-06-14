package main

import (
	"math/rand"
	"rps/api"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	api.Run()
}
