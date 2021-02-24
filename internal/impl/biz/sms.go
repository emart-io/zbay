package biz

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func SendSMS(to, title, address string) {
	// TODO
	to = "+8615901251201"
	// https://cloud.tencent.com/document/product/382/13301
	if len([]rune(title)) > 12 {
		title = title[0:23]
	}
	if len([]rune(address)) > 12 {
		address = address[0:23]
	}
	credential := common.NewCredential(smsSecretId, smsSecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{to})
	request.TemplateParamSet = common.StringPtrs([]string{title, address})
	request.SmsSdkAppid = common.StringPtr(smsAppId)
	request.Sign = common.StringPtr("土产直卖")
	request.TemplateID = common.StringPtr("875567")

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		//panic(err)
		log.Errorln(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
