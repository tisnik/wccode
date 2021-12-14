package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	hash := md5.New()

	io.WriteString(hash, "Příliš žluťoučký kůň")
	fmt.Printf("%x", hash.Sum(nil))
}
