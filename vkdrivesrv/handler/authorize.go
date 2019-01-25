package handler

import (
	_"fmt"
	"net/http"
    "io/ioutil"
    "errors"

    "../vkservice"
    "../db"
    "../config"
    "../crypt"
    
    
	"github.com/bitly/go-simplejson"
)


//Authorize is used to authorize user in application
var Authorize = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
    //traceHTTPRequest(request)

    code := request.URL.Query().Get("code")
    if code == "" {
        responseMessage(GeneralResponse{1, "No code was passed"}, response, request)
    }

    res, err := loginUserIntoVK(code, request.UserAgent())
    if err != nil {
        responseMessage(GeneralResponse{1, err.Error()}, response, request)
    } else {
        responseMessage(res, response, request)
    }
    return
})


func loginUserIntoVK(strCode, strUserAgent string) (LoginResponse, error){

    var loginRespone LoginResponse

    clientSecret := config.Config.App.ClientSecret
    clientId := config.Config.App.ClientId
    redirectUri := config.Config.App.RedirectUri
    
    resp, err := http.Get("https://oauth.vk.com/access_token?client_id=" + clientId + "&client_secret=" + clientSecret + "&redirect_uri=" + redirectUri + "&code=" + strCode)
    if err != nil {
        return loginRespone, errors.New("Can't send request for getting access token")
    }   

    body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
    accessToken, err := json.Get("access_token").String()
    userSNId, _ := json.Get("user_id").Int()

    if err != nil {
        return loginRespone, errors.New("User are not authorize in VK")
    }

    fName, lName, photo, err := vkservice.GetUserInfo(accessToken)
    if err != nil {
        return loginRespone, err
    }

    user, err := db.SelectUserBySNId(userSNId, "vk")
    if err == nil {
        err = db.UpdateUser(user.Id, fName, lName, photo, true)
        if err != nil {
            return loginRespone, err
        }
    } else {
        err = db.InsertUser(userSNId, "vk", fName, lName, photo)
        if err != nil {
            return loginRespone, err
        }
        user, _ = db.SelectUserBySNId(userSNId, "vk")
    }

    sessionToken := crypt.GetSessionToken(user.Id)
    db.InsertSession(user.Id, sessionToken, strUserAgent)

    jwtToken := crypt.JWTSign(crypt.JWTClaim{SNId:user.Id, SNType:"vk", Token: sessionToken})

    loginRespone = LoginResponse{GeneralResponse{0,""}, fName, lName, photo, jwtToken}

    return loginRespone, nil
}