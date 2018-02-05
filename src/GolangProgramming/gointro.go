package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	fmt.Println("It is currently:",time.Now())
	fmt.Println("A number from 1-66", rand.Intn(66))

}