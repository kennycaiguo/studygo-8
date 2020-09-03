package main

import (
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	timestamp := time.Now().Unix()
	main()
	t.Log(timestamp)
}
