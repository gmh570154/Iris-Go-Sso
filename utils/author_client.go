package utils

import (
	"encoding/json"
	"io/ioutil"
	"iris_master/common/configs"
	"iris_master/log"
	"net/http"
	"net/url"
	"strings"
)

const Bearer = "Bearer"

type RequestBody struct {
	grant_type   string
	code         string
	redirect_uri string
}

// Prepare a request
func CodeOauth(code string, redirect_uri string) (result bool) {
	data := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {redirect_uri},
	}

	req_body := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", configs.AppConfig.Sso.HostUrl+"/v1/oauth/tokens", req_body)

	if err != nil {
		return false
	}
	request.SetBasicAuth("test_client_3", "test_secret")

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	log.Log.Info(string(body))
	var jsonMap map[string]interface{}
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		log.Log.Error(err)
		return false
	}
	log.Log.Info(jsonMap)
	//{"user_id":"642dec7d-2bb2-46f5-9d6f-27d12dbf92d2","access_token":"2a6ced51-048b-45b0-a7f7-869a6bd0798f","expires_in":3600,"token_type":"Bearer","scope":"read","refresh_token":"13032f1b-286d-4d9f-831b-7670b0732fae"}
	if strings.Contains(string(body), "user_id") {
		return true
	} else {
		return false
	}
}

// Prepare a request
func TokenOauth(token string, redirect_uri string) (result bool) {
	data := url.Values{
		"token_type_hint": {"access_token"},
		"token":           {token},
		//	"redirect_uri": {redirect_uri},
	}

	req_body := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", configs.AppConfig.Sso.HostUrl+"/v1/oauth/introspect", req_body)

	if err != nil {
		return false
	}
	request.SetBasicAuth("test_client_3", "test_secret")

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	log.Log.Info(string(body))
	var jsonMap map[string]interface{}
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		log.Log.Error(err)
		return false
	}

	log.Log.Info(jsonMap)
	//{"user_id":"642dec7d-2bb2-46f5-9d6f-27d12dbf92d2","access_token":"2a6ced51-048b-45b0-a7f7-869a6bd0798f","expires_in":3600,"token_type":"Bearer","scope":"read","refresh_token":"13032f1b-286d-4d9f-831b-7670b0732fae"}
	if strings.Contains(string(body), "user_id") || strings.Contains(string(body), "active") {
		return true
	} else {
		return false
	}
}
