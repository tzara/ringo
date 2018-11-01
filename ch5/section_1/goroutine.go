package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start", time.Now())
	go long()
	go short()
	time.Sleep(5 * time.Second)
	fmt.Println("main end", time.Now())
}

func long() {
	fmt.Println("long start", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long end", time.Now())
}

func short() {
	fmt.Println("short start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short end", time.Now())
}
