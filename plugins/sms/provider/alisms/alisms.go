package alisms

import (
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 请求接口
const API = "http://dysmsapi.aliyuncs.com/"

type Sender struct {
	Ak string
	Sk string
	// 请求参数
	Params map[string]string
}

// 初始化一些参数
func New(ak string, sk string) *Sender {

	params := make(map[string]string)

	params["AccessKeyId"] = ak
	params["RegionId"] = "cn-hangzhou"
	params["Format"] = "JSON"
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = getNonce()
	// GMT 日期格式 2019-05-08T10:20:40Z
	params["Timestamp"] = getGMTime()
	params["Action"] = "SendSms"
	params["Version"] = "2017-05-25"

	sd := &Sender{}
	sd.Ak = ak
	sd.Sk = sk
	sd.Params = params

	return sd
}

// 发送请求
func (sd *Sender) Request(phoneNumbers string, signName string, tplCode string, tplParams string) (string, error) {

	sd.Params["PhoneNumbers"] = phoneNumbers
	sd.Params["SignName"] = signName
	sd.Params["TemplateCode"] = tplCode
	sd.Params["TemplateParam"] = tplParams
	sd.Params["Signature"] = sd.computeSignature(sd.Params)

	payload := strings.NewReader(getPayload(sd.Params))

	req, _ := http.NewRequest("POST", API, payload)

	req.Header.Set("x-sdk-client", "php/2.0.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// 获取请求body
func getPayload(params map[string]string) (payload string) {

	for k := range params {
		payload += k + "=" + url.QueryEscape(params[k]) + "&"
	}

	return strings.Trim(payload, "&")
}

// 设置为gmt时间
func getGMTime() string {

	return time.Now().UTC().Format("2006-01-02T15:04:05Z")

}

// 计算签名
func (sd *Sender) computeSignature(params map[string]string) string {

	// 对请求参数key进行排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sign string

	for _, k := range keys {
		sign += "&" + strEncode(k) + "=" + strEncode(params[k])
	}

	sign = "POST&%2F&" + strEncode(strings.Trim(sign, "&"))

	// hmac sha1
	h := hmac.New(sha1.New, []byte(sd.Sk+"&"))
	io.WriteString(h, sign)

	return b64.StdEncoding.EncodeToString(h.Sum(nil))

}

func strEncode(str string) string {

	str = url.QueryEscape(str)

	reg, _ := regexp.Compile(`\+`)
	str = reg.ReplaceAllString(str, "%20")

	reg, _ = regexp.Compile(`\*`)
	str = reg.ReplaceAllString(str, "%2A")

	reg, _ = regexp.Compile(`%7E`)
	str = reg.ReplaceAllString(str, "~")

	return str
}

// 设置随机数
func getNonce() string {

	now := time.Now()
	secs := now.Unix()
	return strconv.Itoa(int(secs))
}
