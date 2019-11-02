package biz

import (
	"fmt"
	"testing"
)

func TestUint64(t *testing.T) {
	a := uint64(222)
	i := float32(a) / 100
	fmt.Println(i)
	t.Error("end")
}
