package biz

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"sort"
	"time"

	pb "github.com/emart.io/zbay/service/go"
	log "github.com/sirupsen/logrus"
)

func AlipayRequest(pm *pb.PayMap) (p *pb.PayMap, err error) {
	pm.Kv["app_id"] = alipayAppId
	pm.Kv["charset"] = "utf-8"
	pm.Kv["sign_type"] = "RSA2"
	pm.Kv["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	pm.Kv["version"] = "1.0"
	pm.Kv["sign"], err = sign(pm.Kv)
	if err != nil {
		return nil, err
	}

	return pm, nil
}

// type BizContent struct {
// 	Subject     string  `json:"subject"`
// 	OutTradeNo  string  `json:"out_trade_no"`
// 	TotalAmount float32 `json:"total_amount"`
// 	ProductCode string  `json:"product_code"`
// 	QuitUrl     string  `json:"quit_url"`
// }

// func SignAlipay(totalAmount float32) (string, error) {
// 	bizContent := BizContent{
// 		Subject:     "订单费用",
// 		OutTradeNo:  "bj-gd-" + fmt.Sprint(time.Now().Unix()) + fmt.Sprint(mr.Intn(10)),
// 		ProductCode: "QUICK_MSECURITY_PAY", // QUICK_WAP_WAY | QUICK_MSECURITY_PAY
// 		TotalAmount: totalAmount,
// 	}

// 	data := make(map[string]string)
// 	data["app_id"] = alipayAppId
// 	data["method"] = "alipay.trade.app.pay" // alipay.trade.wap.pay|alipay.trade.app.pay
// 	// if returnUrl != "" {
// 	// 	data.Add("return_url",returnUrl) }
// 	// 	 if notifyUrl != "" {
// 	// 		 data.Add("notify_url", returnUrl) }
// 	//data.Add("format", "json")
// 	data["charset"] = "utf-8"
// 	data["sign_type"] = "RSA2"
// 	data["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
// 	data["version"] = "1.0"
// 	bc, err := json.Marshal(bizContent)
// 	if err != nil {
// 		log.Errorln(err)
// 		return "", err
// 	}
// 	data["biz_content"] = string(bc)
// 	data["sign"], _ = Sign(data)
// 	return "", nil //data.Encode(), nil
// }

func sign(m map[string]string) (string, error) {
	//对url.values进行排序
	sign := ""
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if m[k] != "" {
			if i == 0 {
				sign = k + "=" + m[k]
			} else {
				sign = sign + "&" + k + "=" + m[k]
			}
		}
	}
	b, err := rsaEncrypt([]byte(sign))
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	log.Infoln("加密：", b)
	log.Infoln("base加密：", base64.StdEncoding.EncodeToString(b))
	return base64.StdEncoding.EncodeToString(b), nil
}

func rsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(alipayPrivateKey))
	if block == nil {
		log.Errorln("block空")
		return nil, errors.New("certificate failed to load")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Errorln("无法还原私钥", err)
		return nil, err
	}
	h2 := sha256.New()
	h2.Write(origData)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, h2.Sum(nil))
}
