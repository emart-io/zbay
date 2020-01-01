package biz

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	pb "github.com/emart.io/zbay/service/go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Body             string `xml:"body"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Trade_type       string `xml:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no"`
	Sign             string `xml:"sign"`
}

func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign

}

func WechatPayParams(order *pb.Order) (*pb.WechatPayParams, error) {
	var yourReq UnifyOrderReq
	yourReq.Appid = wechatAppId
	yourReq.Body = "订单费用"
	yourReq.Mch_id = wechatMchId
	yourReq.Nonce_str = fmt.Sprint(time.Now().Unix())
	yourReq.Notify_url = "www.yourserver.com/wxpayNotify"
	yourReq.Trade_type = "APP"
	yourReq.Spbill_create_ip = "10.0.0.1"
	yourReq.Total_fee = int(order.Amount) //单位是分
	yourReq.Out_trade_no = fmt.Sprint(time.Now().UnixNano()) + fmt.Sprint(rand.Intn(10))

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["body"] = yourReq.Body
	m["mch_id"] = yourReq.Mch_id
	m["notify_url"] = yourReq.Notify_url
	m["trade_type"] = yourReq.Trade_type
	m["spbill_create_ip"] = yourReq.Spbill_create_ip
	m["total_fee"] = yourReq.Total_fee
	m["out_trade_no"] = yourReq.Out_trade_no
	m["nonce_str"] = yourReq.Nonce_str

	yourReq.Sign = wxpayCalcSign(m, wechatKey) //这个是计算wxpay签名的函数上面已贴出
	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return nil, err
	}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{Timeout: 5 * time.Second}
	resp, _err := c.Do(req)
	if _err != nil {
		fmt.Println("请求微信支付统一下单接口发送错误, 原因:", _err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取http body失败，原因!", err)
		//http.Error(w.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, err
	}

	defer resp.Body.Close()
	xmlResp := UnifyOrderResp{}
	//body := resp.Body
	_err = xml.Unmarshal(body, &xmlResp)
	//处理return code.
	if xmlResp.Return_code == "FAIL" {
		fmt.Println("微信支付统一下单不成功，原因:", xmlResp.Return_msg)
		return nil, err
	}

	//这里已经得到微信支付的prepay id，需要返给客户端，由客户端继续完成支付流程
	fmt.Println("微信支付统一下单成功，预支付单号:", xmlResp.Prepay_id)
	var ns map[string]interface{}
	ns = make(map[string]interface{}, 0)
	ns["appid"] = yourReq.Appid
	ns["partnerid"] = yourReq.Mch_id
	ns["prepayid"] = xmlResp.Prepay_id
	ns["package"] = "Sign=WXPay"
	ns["noncestr"] = "nonstr" + fmt.Sprint(time.Now().Unix())
	ns["timestamp"] = fmt.Sprint(time.Now().Unix())
	ns["sign"] = wxpayCalcSign(ns, wechatKey)

	return &pb.WechatPayParams{
		Partnerid: ns["partnerid"].(string),
		Prepayid:  ns["prepayid"].(string),
		Noncestr:  ns["noncestr"].(string),
		Timestamp: ns["timestamp"].(string),
		Sign:      ns["sign"].(string),
	}, nil

}

type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
}
