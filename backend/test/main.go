package main

import (
	"fmt"
	"time"
)

func main() {
	exp := time.Now().Add(time.Duration(10) * time.Second).Unix()

	time.Sleep(time.Second * time.Duration(11))

	fmt.Printf("time.Now().Before(time.Unix(exp, 0)): %v\n", time.Now().Before(time.Unix(exp, 0)))

}
