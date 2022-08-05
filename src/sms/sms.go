package sms

import (
	"errors"
	"log"
	"os"
	"sync"
	"taurus-backend/constant"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	terrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

var client *Client
var once *sync.Once

func GetSMSClient() *Client {
	once.Do(func() {
		client = NewSMSClient()
	})
	return client
}

type Client struct {
	session    *sms.Client
	credential *common.Credential
	profile    *profile.ClientProfile
}

func CheckSmsEnv() {
	if os.Getenv("SECRET_ID") == "" ||
		os.Getenv("SECRET_KEY") == "" ||
		os.Getenv("SMS_SDK_APP_ID") == "" ||
		os.Getenv("SMS_SIGN_NAME") == "" ||
		os.Getenv("TEMPLATE_ID") == "" {
		log.Fatal("SMS env not correct.")
	}
}

func NewSMSClient() *Client {
	c := &Client{}
	return c
}

func (c *Client) Init(e *constant.Env) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.NetworkFailureMaxRetries = 3
	cpf.NetworkFailureRetryDuration = profile.ConstantDurationFunc(5)

	credential := common.NewCredential(e.SecretId, e.SecretKey)
	client, err := sms.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		log.Fatalln("NewSMSClient error:", err)
		return
	}
	c.profile = cpf
	c.session = client
	c.credential = credential
}

func (c *Client) SendSMS(phone string, awardType int, awardCode string) (smsSerialNo string, err error) {
	req := sms.NewSendSmsRequest()
	req.PhoneNumberSet = common.StringPtrs([]string{phone})
	req.SmsSdkAppId = common.StringPtr(os.Getenv("SMS_SDK_APP_ID"))
	req.SignName = common.StringPtr(os.Getenv("SMS_SIGN_NAME"))
	req.TemplateId = common.StringPtr(os.Getenv("TEMPLATE_ID"))
	req.TemplateParamSet = getTemplateParamSet(awardType, awardCode)

	resp, err := c.session.SendSms(req)
	if _, ok := err.(*terrors.TencentCloudSDKError); ok {
		log.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		return "", err
	}
	sendStatus := resp.Response.SendStatusSet[0]
	if *sendStatus.Code != "Ok" || *resp.Response.SendStatusSet[0].SerialNo == "" {
		return "", errors.New("send sms fail, code is not ok or serial no is empty")
	}
	return *resp.Response.SendStatusSet[0].SerialNo, nil
}

func getTemplateParamSet(awardType int, code string) []*string {
	var paramSet []*string
	switch awardType {
	case constant.MEITUAN:
		paramSet = common.StringPtrs([]string{"美团外卖五元券", code, "美团APP"})
	case constant.TENCENT:
		paramSet = common.StringPtrs([]string{"腾讯视频月卡", code, "腾讯视频App"})
	case constant.DIDI:
		paramSet = common.StringPtrs([]string{"青桔单车周卡", code, "青桔单车APP/小程序"})
	default:
	}
	return paramSet
}
