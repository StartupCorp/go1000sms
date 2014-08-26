package go1000sms

import (
	"errors"
	"gocms/app/go1000sms/go1000smstools"
)

const (
	SmsApiRoot = "http://api.1000sms.ru/"
)

//==============================================================================
// Go1000sms
//================================================
type Go1000sms struct {
	Email      string
	Password   string
	SenderName string
}

//==============================================================================
func New(email, password, sender_name string) *Go1000sms {
	return &Go1000sms{Email: email, Password: password, SenderName: sender_name}
}

//==============================================================================
func (go1000sms *Go1000sms) PushMsg(phone, text string) error {
	data := make(map[string]string)
	data["method"] = "push_msg"
	data["format"] = "json"
	data["email"] = go1000sms.Email
	data["password"] = go1000sms.Password
	data["sender_name"] = go1000sms.SenderName
	data["phone"] = phone
	data["text"] = text

	resp, err := go1000smstools.Post(SmsApiRoot, data)
	if err != nil {
		return err
	}

	msg := resp["response"].(map[string]interface{})["msg"].(map[string]interface{})
	if msg["err_code"] != "0" {
		return errors.New(msg["text"].(string))
	}

	return nil
}

//==============================================================================
func (go1000sms *Go1000sms) GetProfile() (map[string]interface{}, error) {
	data := make(map[string]string)
	data["method"] = "get_profile"
	data["format"] = "json"
	data["email"] = go1000sms.Email
	data["password"] = go1000sms.Password

	resp, err := go1000smstools.Post(SmsApiRoot, data)
	if err != nil {
		return nil, err
	}

	msg := resp["response"].(map[string]interface{})["msg"].(map[string]interface{})
	if msg["err_code"] != "0" {
		return nil, errors.New(msg["text"].(string))
	}

	return resp["response"].(map[string]interface{})["data"].(map[string]interface{}), nil
}

//==============================================================================
