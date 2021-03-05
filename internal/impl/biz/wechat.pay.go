package biz

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	pb "github.com/emart.io/zbay/service/go"
)

// UnifiedRequest ... 统一下单
func UnifiedRequest(pm *pb.PayMap, ip string) (*pb.PayMap, error) {
	pm.Kv["appid"] = wechatAppId
	pm.Kv["mch_id"] = wechatMchId
	pm.Kv["nonce_str"] = fmt.Sprint(time.Now().Unix())
	pm.Kv["spbill_create_ip"] = ip

	pm.Kv["sign"] = WXpayCalcSign(pm.Kv)

	bytes_req, err := xml.Marshal(StringMap(pm.Kv))
	if err != nil {
		fmt.Println("以xml形式编码发送错误, 原因:", err)
		return nil, err
	}

	str_req := string(bytes_req)
	fmt.Println(str_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "StringMap", "xml", -1)
	bytes_req = []byte(str_req)

	//发送unified order请求.
	req, err := http.NewRequest("POST", pm.Url, bytes.NewReader(bytes_req))
	if err != nil {
		fmt.Println("New Http Request发生错误，原因:", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/xml")
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	resp, _err := (&http.Client{Timeout: 5 * time.Second}).Do(req)
	if _err != nil {
		fmt.Printf("请求微信支付接口%s发送错误, 原因:%s", pm.Url, _err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取http body失败，原因!", err)
		//http.Error(w.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, err
	}

	defer resp.Body.Close()
	xmlResp := pb.PayMap{Kv: make(map[string]string)}
	str_req = strings.Replace(string(body), "xml", "StringMap", -1)
	//body := resp.Body
	_err = xml.Unmarshal([]byte(strings.Replace(string(body), "xml", "StringMap", -1)), (*StringMap)(&xmlResp.Kv))
	//处理return code.
	if xmlResp.Kv["return_code"] == "FAIL" {
		fmt.Printf("微信支付接口%s不成功, 原因:%s", pm.Url, xmlResp.Kv["return_msg"])
		return nil, err
	}
	fmt.Println("成功:", xmlResp.Kv)
	return &xmlResp, nil
}

func WXpayCalcSign(mReq map[string]string) (sign string) {
	fmt.Println("微信支付签名计算, API KEY:", wechatKey)
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
	if wechatKey != "" {
		signStrings = signStrings + "key=" + wechatKey
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
