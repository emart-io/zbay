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

func TestWechat(t *testing.T) {
	paras, err := WechatPayParams(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(paras)
	//fmt.Println(i)
	t.Error("end")
}
