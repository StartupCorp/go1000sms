package go1000smstools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
)

//==============================================================================
func Post(myurl string, args map[string]string) (map[string]interface{}, error) {
	val := url.Values{}
	var first bool
	for k, v := range args {
		if !first {
			val.Set(k, v)
			first = true
		} else {
			val.Add(k, v)
		}
	}
	
	fmt.Printf("go1000smstools.Post(): %v, %v\n", myurl, val.Encode())

	resp, err := http.Post(myurl, "application/x-www-form-urlencoded", bytes.NewBufferString(val.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//================================================
func Get(myurl string) (map[string]interface{}, error) {
	fmt.Printf("go1000smstools.Get(): %v\n", myurl)

	resp, err := http.Get(myurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//==============================================================================
