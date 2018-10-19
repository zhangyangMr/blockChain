package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	data := "hello world"
	for i := 0; i < 1000000; i++{
		hash := sha256.Sum256([]byte(data + string(i)))
		fmt.Printf("hash : %x \n",hash[:])
	}
}
