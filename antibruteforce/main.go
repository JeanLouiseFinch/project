package main

import (
	"antibruteforce/cmd"
	"log"
	"time"
)

func main() {
	<-time.After(30 * time.Second)
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
