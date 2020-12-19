package bfs

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
