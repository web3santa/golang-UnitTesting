package math

import (
	"log"
	"os"
)

func Abs(a int) int {
	path := os.Getenv("GOPATH")
	log.Println(path)
	if a > 0 {
		return a
	}

	return -a
}
