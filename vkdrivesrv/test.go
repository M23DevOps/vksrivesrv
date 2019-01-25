package main

import (
	_"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"

	_"./handler"
	"./config"
	"./db"
	_"./vkservice"
	_"./model"
	"./crypt"

	"github.com/bitly/go-simplejson"
	"github.com/karlseguin/typed"
)


func main() {

	config.LoadConfiguration("./config.json")

	Test6() 
}

func Test6() {
	db.ConnectToDB()
	lol, err := db.SelectImagesByUser(2)
	fmt.Println(err)
	fmt.Println(lol)
}

func Test5() {
	parts := strings.Split("lol.kek", ".")
	fmt.Println(len(parts))
}

func Test4() {
	ff := crypt.JWTClaim{SNId:1, SNType : "f", Token:"f"}
	fmt.Println(ff)

	tt, _ := json.Marshal(ff)

	var hh crypt.JWTClaim
	err := json.Unmarshal(tt,&hh)

	fmt.Println(hh)
	fmt.Println(err)
}

func Test3() {
	//err := db.InsertSession(1,"token","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	//_,err := db.SelectSession("token","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	err :=db.DeleteSession("token","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	
	fmt.Println(err)
}

func Test2() {

	claim := crypt.JWTClaim{Token:"lol"}

	token := crypt.JWTSign(claim)
	fmt.Println(token)


	fmt.Println(crypt.GetSessionToken(1))
}
	
func Test1() {
	resp, _ := http.Get("https://api.vk.com/method/messages.getDialogs?count=3&access_token=a9d58a6fe744f8171ea19b9ca95343009e0938920f717467046bb07bfb47173fc8b7047306b0388f412d1&v=5.60")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	sj, _ := simplejson.NewJson(body)

	_, err := sj.Get("response").Get("items").GetIndex(0).Get("message").Get("idfff").String()

	fmt.Println(err)

	tt,_ := typed.Json(body)

	lol := tt.Object("response").Objects("items")[0].Object("message").Int("id")

	fmt.Println(lol)
}
