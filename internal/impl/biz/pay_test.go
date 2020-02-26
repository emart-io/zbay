package biz

import (
	"fmt"
	"testing"

	pb "github.com/emart.io/zbay/service/go"
)

func TestUint64(t *testing.T) {
	a := uint64(222)
	i := float32(a) / 100
	fmt.Println(i)
	t.Error("end")
}

func TestWechatPrepay(t *testing.T) {
	var pm = pb.PayMap{Kv: make(map[string]string)}
	pm.Kv["body"] = "订单费用"
	pm.Kv["notify_url"] = "www.yourserver.com/wxpayNotify"
	pm.Kv["trade_type"] = "MWEB"
	pm.Kv["total_fee"] = "1"
	pm.Kv["out_trade_no"] = "casdfasdfe"

	pm.Url = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	paras, err := PayRequest(&pm, "192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(paras)
	//fmt.Println(i)
	t.Error("end")
}

func TestWechatQuery(t *testing.T) {
	var pm = pb.PayMap{Kv: make(map[string]string)}
	pm.Kv["out_trade_no"] = "casdfasdfe"

	pm.Url = "https://api.mch.weixin.qq.com/pay/orderquery"
	paras, err := PayRequest(&pm, "192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(paras)
	//fmt.Println(i)
	t.Error("end")
}
