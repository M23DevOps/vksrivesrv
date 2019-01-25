package vkservice

import (
	_"fmt"
	"io/ioutil"
	"net/http"
	"errors"

	"github.com/bitly/go-simplejson"
)


func GetUserInfo(strToken string) (string, string, string, error){
	resp, err := http.Get("https://api.vk.com/method/users.get?fields=photo_50&access_token=" + strToken + "&v=5.67")
	if err != nil {
		return "","","", errors.New("Can't get user info")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
	f_name, _ := json.Get("response").GetIndex(0).Get("first_name").String()
	l_name, _ := json.Get("response").GetIndex(0).Get("last_name").String()
	photo, _ := json.Get("response").GetIndex(0).Get("photo_50").String()

	return f_name, l_name, photo, nil
}
