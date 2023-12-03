package main

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(time.Millisecond * 200) // Needed to let server spin up

	m.Run()

	os.Interrupt.Signal()
}
