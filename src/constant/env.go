package constant

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Stage       string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	SecretId    string
	SecretKey   string
	SmsSdkAppId string
	SmsSignName string
	TemplateId  string
	Multiple    string
}

func (e *Env) Check() {
	if "" == e.DBHost ||
		"" == e.DBPort ||
		"" == e.DBUser ||
		"" == e.DBPassword ||
		"" == e.DBName ||
		"" == e.SecretId ||
		"" == e.SecretKey ||
		"" == e.SmsSdkAppId ||
		"" == e.SmsSignName ||
		"" == e.TemplateId {
		log.Fatalf("missing require env: %v", e)
	}

}

func (e *Env) Init() {
	if "" == os.Getenv("STAGE") {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("get local env err:", err)
		}
	}
	e.Stage = os.Getenv("STAGE")
	e.DBHost = os.Getenv("DB_HOST")
	e.DBPort = os.Getenv("DB_PORT")
	e.DBUser = os.Getenv("DB_USER")
	e.DBPassword = os.Getenv("DB_PASSWORD")
	e.DBName = os.Getenv("DB_NAME")
	e.SecretId = os.Getenv("SECRET_ID")
	e.SecretKey = os.Getenv("SECRET_KEY")
	e.SmsSdkAppId = os.Getenv("SMS_SDK_APP_ID")
	e.SmsSignName = os.Getenv("SMS_SIGN_NAME")
	e.TemplateId = os.Getenv("TEMPLATE_ID")
	e.Multiple = os.Getenv("MULTIPLE")
	if "" == e.Multiple {
		e.Multiple = "1"
	}
}

func (e *Env) GetDBConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", e.DBUser, e.DBPassword, e.DBHost, e.DBPort, e.DBName)
}
