package main

import (
	_ "go.uber.org/automaxprocs"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the service
	// Based on what is available either by the matches or quotes.
	//if _, err := maxprocs.Set(); err != nil {
	//	fmt.Printf("maxprocs: %s \n", err)
	//	os.Exit(1)
	//}

	g := runtime.GOMAXPROCS(0)
	log.Printf("starting service build[%s] CPU[%d]", build, g)
	defer log.Println("service ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping service")
}

//https://github.com/ardanlabs/service
//https://github.com/ardanlabs/service/tree/service3
