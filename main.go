package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	deaths := Conduct(100)

	fmt.Println()
	fmt.Println("Total Deaths: " + strconv.Itoa(deaths))

	fmt.Println()
	fmt.Println("Trials have successfully been conducted, press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
