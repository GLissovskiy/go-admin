package main

import (
	"testing"
	"time"
)

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)

}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestParallelThree(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
